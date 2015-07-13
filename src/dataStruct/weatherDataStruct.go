package dataStruct

import (

)

type Data struct{
	Status string
	Weather [] Weatherdata
}

type Weatherdata struct{
	City_name string 
	City_id string
	Last_update string
	Now NowWeather
	Today Todaydata
	Future [] Futuredata
}

type NowWeather struct{
	Text string
	Code string
	Temperature string
	Feels_like string
	Wind_direction string
	Wind_speed string
	Wind_scale string 
	Humidity string
	Visibility string
	Pressure string 
	Pressure_rising string
	Air_quality string 
}

type Todaydata struct{
	Sunrise string 
	Sunset string
	Suggestion string
}

type Futuredata struct{
	Date string
	Day string
	Text string
	Code1 string
	Code2 string
	High string
	Low string
	Cop string
	Wind string
}