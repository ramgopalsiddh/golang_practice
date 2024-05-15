# Weather web-app
- This Project use [Zomato Weather Union's API](https://www.weatherunion.com) 
- That use latitude & longitude for get data from [Weather Union's API](https://www.weatherunion.com) 
- I created localities database table that store city name, locality name, latitude, longitude & Device type.
- Localities table is used to map latitude, longitude by give city & locality name.


### notes
- Project needs Zomato Weather Union's api key under `WEATHERUNION_API_KEY` env var.
- you can set the API key with `export WEATHERUNION_API_KEY=<your_api_key>`
- Get the API Key and API docs from https://www.weatherunion.com/dashboard/
- Create localities table by use localities.csv
- Then run main.go 
- visit [localhost8080](http://localhost:8080/)


