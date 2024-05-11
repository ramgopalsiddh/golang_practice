package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

const (
    baseURL   = "https://www.weatherunion.com/gw/weather/external/v0"
    apiKey    = "api key"
)

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

func weatherHandler(w http.ResponseWriter, r *http.Request) {
    latitudeStr := r.URL.Query().Get("latitude")
    longitudeStr := r.URL.Query().Get("longitude")

    if latitudeStr == "" || longitudeStr == "" {
        http.Error(w, "Latitude and longitude are required", http.StatusBadRequest)
        return
    }

    latitude, err := strconv.ParseFloat(latitudeStr, 64)
    if err != nil || latitude < -90 || latitude > 90 {
        http.Error(w, "Invalid latitude", http.StatusBadRequest)
        return
    }

    longitude, err := strconv.ParseFloat(longitudeStr, 64)
    if err != nil || longitude < -180 || longitude > 180 {
        http.Error(w, "Invalid longitude", http.StatusBadRequest)
        return
    }

    // Set up request to Weather Union API
    url := fmt.Sprintf("%s/get_weather_data?latitude=%s&longitude=%s", baseURL, latitudeStr, longitudeStr)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }

    // Set the required headers
    req.Header.Set("content-type", "application/json")
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

func main() {
    http.HandleFunc("/get_weather_data", weatherHandler)
    http.Handle("/", http.FileServer(http.Dir("./frontend")))
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
