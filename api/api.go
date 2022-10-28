package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gross2001/weather/model"
)

var apiKey = "b7696781ba8d325ebc82c5f609f05210"
var limit = "5"

func GetCoordinates(cityName string) model.Cities {

	URL := "http://api.openweathermap.org/geo/1.0/direct?q=" + cityName + "&limit=" + limit + "&appid=" + apiKey

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	var cities model.Cities
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(cities)
	return cities
}

func GetWeather(city model.City) model.WeatherRs {
	lat := fmt.Sprintf("%g", city.Lat)
	lon := fmt.Sprintf("%g", city.Lon)

	URL := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + apiKey

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	var weather model.WeatherRs
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Fatal(err.Error())
	}
	return weather
}
