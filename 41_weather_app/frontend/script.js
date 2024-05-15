document.addEventListener("DOMContentLoaded", function() {
    fetchCities();
});

function fetchCities() {
    fetch("/get_cities")
        .then(response => response.json())
        .then(data => {
            const citySelect = document.getElementById("city");
            data.forEach(city => {
                const option = document.createElement("option");
                option.value = city.city_name;
                option.textContent = city.city_name;
                citySelect.appendChild(option);
            });
            // Trigger fetchLocalities() when the city select changes
            citySelect.addEventListener("change", fetchLocalities);
        })
        .catch(error => console.error("Error fetching cities:", error));
}

function fetchLocalities() {
    const city = document.getElementById("city").value;
    fetch(`/get_localities?city=${encodeURIComponent(city)}`)
        .then(response => response.json())
        .then(data => {
            const localitySelect = document.getElementById("locality");
            // Clear existing options
            localitySelect.innerHTML = "";
            data.forEach(locality => {
                const option = document.createElement("option");
                option.value = locality.locality_name;
                option.textContent = locality.locality_name;
                localitySelect.appendChild(option);
            });
        })
        .catch(error => console.error("Error fetching localities:", error));
}

document.getElementById("submit-btn").addEventListener("click", function() {
    getWeather();
});

function getWeather() {
    const city = document.getElementById("city").value;
    const locality = document.getElementById("locality").value;
    
    // Fetch weather data from backend
    fetch(`/get_weather_data?city=${encodeURIComponent(city)}&locality=${encodeURIComponent(locality)}`)
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to fetch weather data");
            }
            return response.json();
        })
        .then(weatherData => {
            // Convert wind direction from degrees to direction names
            const windDirectionName = degToDir(weatherData.locality_weather_data.wind_direction);

            // Check if device type is 2 (rain gauge)
            if (weatherData.device_type === 2) {
                document.getElementById("weather-data").innerHTML = `
                    <p><strong style="color:red">${weatherData.message}</strong></p>
                    <p>This location has <strong>Rain gauge system</strong> that measures rain only.</p>
                    <p><strong>Temperature &#x1F321:</strong> ${weatherData.locality_weather_data.temperature} °C;</p>
                    <p><strong>Humidity &#x1F4A7:</strong> ${weatherData.locality_weather_data.humidity} % ;</p>
                    <p><strong>Wind Speed &#x1F343:</strong> ${weatherData.locality_weather_data.wind_speed} m/s;</p>
                    <p><strong>Wind Direction &#x1F4A8:</strong>${windDirectionName}/ ${weatherData.locality_weather_data.wind_direction}°;</p>
                    <p><strong>Rain Intensity &#x1F327:</strong> ${weatherData.locality_weather_data.rain_intensity} mm/hr;</p>
                    <p><strong>Rain Accumulation &#x1F327:</strong> ${weatherData.locality_weather_data.rain_accumulation} mm ;</p>
                `;
                return;
            }

            // Update the weather data in the frontend
            const weatherDisplay = document.getElementById("weather-data");
            weatherDisplay.innerHTML = `
                <p><strong style="color:red">${weatherData.message}</strong></p>
                <p>This location has <strong>Automated weather system</strong><p>
                <p><strong>Temperature &#x1F321:</strong> ${weatherData.locality_weather_data.temperature} °C;</p>
                <p><strong>Humidity &#x1F4A7:</strong> ${weatherData.locality_weather_data.humidity} % ;</p>
                <p><strong>Wind Speed &#x1F343:</strong> ${weatherData.locality_weather_data.wind_speed} m/s;</p>
                <p><strong>Wind Direction &#x1F4A8:</strong> ${windDirectionName} / ${weatherData.locality_weather_data.wind_direction}°;</p>
                <p><strong>Rain Intensity &#x1F327:</strong> ${weatherData.locality_weather_data.rain_intensity} mm/hr;</p>
                <p><strong>Rain Accumulation &#x1F327:</strong> ${weatherData.locality_weather_data.rain_accumulation} mm ;</p>
            `;
            document.getElementById("error-msg").textContent = "";
        })
        .catch(error => {
            console.error("Error fetching weather data:", error);
            document.getElementById("error-msg").textContent = "Failed to fetch weather data";
        });
}

// Function to convert wind direction from degrees to direction names
function degToDir(deg) {
    const directions = ["N", "NE", "E", "SE", "S", "SW", "W", "NW"];
    const index = Math.round(deg / 45) % 8;
    return directions[index];
}