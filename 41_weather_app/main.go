package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

const (
	baseURL = "https://www.weatherunion.com/gw/weather/external/v0"
	dbUser  = "ram"
	dbName  = "weatherunion"
	dbHost  = "localhost"
	dbPort  = 5432
)
var apiKey string

type City struct {
	CityName string `json:"city_name"`
}

type Locality struct {
	LocalityName string `json:"locality_name"`
	Latitude     float64
	Longitude    float64
}

type Coordinates struct {
    Latitude     float64
	Longitude    float64
}

type WeatherData struct {
	Status             string           `json:"status"`
	Message            string           `json:"message"`
	DeviceType         int              `json:"device_type"`
	LocalityWeatherData LocalityWeather `json:"locality_weather_data"`
}

type LocalityWeather struct {
	Temperature      float64 `json:"temperature"`
	Humidity         float64 `json:"humidity"`
	WindSpeed        float64 `json:"wind_speed"`
	WindDirection    float64 `json:"wind_direction"`
	RainIntensity    float64 `json:"rain_intensity"`
	RainAccumulation float64 `json:"rain_accumulation"`
}

func cityHandler(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	connStr := fmt.Sprintf("user=%s dbname=%s host=%s port=%d sslmode=disable", dbUser, dbName, dbHost, dbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	// Query the database to get city names
	rows, err := db.Query("SELECT DISTINCT city_name FROM localities")
	if err != nil {
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		log.Printf("Failed to query database: %v\n", err)
		return
	}
	defer rows.Close()

	// Fetch city names from the database
	var cities []City
	for rows.Next() {
		var city City
		if err := rows.Scan(&city.CityName); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			log.Printf("Failed to scan row: %v\n", err)
			return
		}
		cities = append(cities, city)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Error retrieving data from database", http.StatusInternalServerError)
		log.Printf("Error retrieving data from database: %v\n", err)
		return
	}

	// Encode and send the data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cities)
}

func localityHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

    // FIXME: clean and uri unescape the city??

	if city == "" {
		http.Error(w, "City is required", http.StatusBadRequest)
		return
	}

	// Connect to the database
	connStr := fmt.Sprintf("user=%s dbname=%s host=%s port=%d sslmode=disable", dbUser, dbName, dbHost, dbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

    // FIXME: fix SQL Injection
	// Query the database to get localities for the city
	rows, err := db.Query("SELECT locality_name FROM localities WHERE city_name = $1", city)
	if err != nil {
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		log.Printf("Failed to query database: %v\n", err)
		return
	}
	defer rows.Close()

	// Fetch localities from the database
	var localities []Locality
	for rows.Next() {
		var locality Locality
		if err := rows.Scan(&locality.LocalityName); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			log.Printf("Failed to scan row: %v\n", err)
			return
		}
		localities = append(localities, locality)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Error retrieving data from database", http.StatusInternalServerError)
		log.Printf("Error retrieving data from database: %v\n", err)
		return
	}

	// Encode and send the data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(localities)
}

func fetchCoordinates(city, locality string) (Coordinates, error) {
    // clean and unescape the city and locality
	// Decode URI parameters
	decodedCity, err := url.QueryUnescape(city)
	if err != nil {
		return Coordinates{}, fmt.Errorf("failed to decode city parameter: %v", err)
	}

	decodedLocality, err := url.QueryUnescape(locality)
	if err != nil {
		return Coordinates{}, fmt.Errorf("failed to decode locality parameter: %v", err)
	}

	// Connect to the database
	connStr := fmt.Sprintf("user=%s dbname=%s host=%s port=%d sslmode=disable", dbUser, dbName, dbHost, dbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to connect to database", http.StatusInternalServerError)
		return Coordinates{}, err
	}
	defer db.Close()

	// Query the database to get latitude and longitude

    // FIXME: check we have sql injection and fix it??
    query := "SELECT latitude, longitude FROM localities WHERE city_name=$1 AND locality_name=$2"
    log.Println("db query: ", query)
	var latitude, longitude float64
	err = db.QueryRow(query, decodedCity, decodedLocality).Scan(&latitude, &longitude)
	if err != nil {
		log.Println("Location not found in database", http.StatusNotFound)
        return Coordinates{}, err
	}

    return Coordinates{Latitude: latitude, Longitude: longitude,}, nil
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	locality := r.URL.Query().Get("locality")

	if city == "" || locality == "" {
		http.Error(w, "City and locality are required", http.StatusBadRequest)
		return
	}

	// Fetch coordinates
	coordinates, err := fetchCoordinates(city, locality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set up request to Weather Union API
	url := fmt.Sprintf("%s/get_weather_data?latitude=%f&longitude=%f", baseURL, coordinates.Latitude, coordinates.Longitude)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	// Set the required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-zomato-api-key", apiKey)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode the response
	var weatherData WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		http.Error(w, "Failed to decode weather data", http.StatusInternalServerError)
		return
	}

	// Set response headers and encode the data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}


func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/about.html")
}

func main() {
    // load API KEY and set the apiKey var
    var found bool
    apiKey, found = os.LookupEnv("WEATHERUNION_API_KEY")
    if !found {
        log.Fatal("WEATHERUNION_API_KEY not found in env variables")
    }
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/get_cities", cityHandler)
	http.HandleFunc("/get_localities", localityHandler)
	http.HandleFunc("/get_weather_data", weatherHandler)
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
    log.Println("Starting server on port 8080...")
    err := http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
	log.Fatal(fmt.Printf("Failed to start the server, %v", err))
}
