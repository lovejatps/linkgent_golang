package main

import (
	"music"
	"net/http"
	"weather"
)

func main() {

	http.HandleFunc("/weather", weather.DefaultHandler)
	http.HandleFunc("/music", music.MusicHandler)
	http.ListenAndServe(":8080", nil)
}
