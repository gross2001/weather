package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gross2001/weather/api"
)

func main() {

	cityName := flag.String("city", "Berlin", "put the name of the city, for example -city London")
	flag.Parse()

	cities, err := api.GetCoordinates(*cityName)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, city := range cities {
		weather, err := api.GetWeather(city)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("for city with lon", weather.Coord.Lon, "and lat", weather.Coord.Lat, "temp", weather.Main.Temp)
	}
}
