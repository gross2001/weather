package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gross2001/weather/model"
)

var (
	limit  = "5"
	domain = "http://api.openweathermap.org"
)

func getToken() (string, error) {

	if val, ok := os.LookupEnv("WEATHER_TOKEN"); ok {
		return val, nil
	}
	err := errors.New("You should provide a WEATHER_TOKEN in enviromental variables")
	return "", err
}

func GetCoordinates(cityName string) (model.Cities, error) {

	var cities model.Cities

	apiKey, err := getToken()
	if err != nil {
		return cities, err
	}

	URL := domain + "/geo/1.0/direct?q=" + cityName + "&limit=" + limit + "&appid=" + apiKey

	resp, err := http.Get(URL)
	if err != nil {
		return cities, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return cities, err
	}
	return cities, nil
}

func GetWeather(city model.City) (model.WeatherRs, error) {

	var weather model.WeatherRs

	lat := fmt.Sprintf("%g", city.Lat)
	lon := fmt.Sprintf("%g", city.Lon)

	apiKey, err := getToken()
	if err != nil {
		return weather, err
	}

	URL := domain + "/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + apiKey
	resp, err := http.Get(URL)
	if err != nil {
		return weather, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return weather, err
	}
	return weather, nil
}
