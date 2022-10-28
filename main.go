package main

import (
	"flag"
	"fmt"

	"github.com/gross2001/weather/api"
)

func main() {

	cityName := flag.String("city", "Berlin", "put the name of the city, for example -city London")
	flag.Parse()

	cities := api.GetCoordinates(*cityName)
	for _, city := range cities {
		weather := api.GetWeather(city)
		fmt.Println("for city with lon", weather.Coord.Lon, "and lat", weather.Coord.Lat, "temp", weather.Main.Temp)
	}
}
