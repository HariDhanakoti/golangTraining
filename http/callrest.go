package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func main() {
	resp,_ := http.Get("https://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=b6907d289e10d714a6e88b30761fae22")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	weatherData := new (WeatherData)
	json.Unmarshal([]byte(body), &weatherData)

	fmt.Println(weatherData)
	fmt.Println(weatherData.MainData.Temp)
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID int
	Main string `json:"main"`
	Desc string `json:"description"`
	Icon string `json:"icon"`
}

type MainData struct {
	Temp float64 `json:"temp"`
	Pressure int64 `json:"pressure"`
	Humidity int64 `json:"humidity"`
	MinTemp float64 `json:"temp_min"`
	MaxTemp float64 `json:"temp_max"`
}

type WeatherData struct {
	Coord Coord
	Weather Weather
	Base string `json:"base"`
	MainData MainData `json:"main"`
	Visibility int64 `json:"visibility"`
}

/*
{
	"coord":
	{
	"lon":-0.13,"lat":51.51
	},
	"weather":[
		{"id":300,
		"main":"Drizzle",
		"description":"light intensity drizzle",
		"icon":"09d"
	}],
	"base":"stations",
	"main":{
		"temp":280.32,
		"pressure":1012,
		"humidity":81,
		"temp_min":279.15,
		"temp_max":281.15
	},
	"visibility":10000,
	"wind":{
		"speed":4.1,
		"deg":80
	},
	"clouds":{
		"all":90
	},
	"dt":1485789600,
	"sys":{
		"type":1,
		"id":5091,
		"message":0.0103,
		"country":"GB",
		"sunrise":1485762037,
		"sunset":1485794875
	},
	"id":2643743,
	"name":"London",
	"cod":200
}

*/