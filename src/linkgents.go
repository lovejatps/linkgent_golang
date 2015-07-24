package main

import (
	"music"
	"net/http"
	"weather"
)

func main() {
	http.HandleFunc("/weather", weather.DefaultHandler)
	http.HandleFunc("/music", music.MusicHandler)    // http://192.168.12.8:8080/music?jsondata=%7B%22artisName%22%3A%22%22%2C%22special%22%3A%22%22%2C%22name%22%3A%22%E7%AC%AC%E4%B8%80%E5%A4%AB%E4%BA%BA%22%2C%22all%22%3A%22%22%2C%22userkey%22%3A%22huxn%22%7D
	http.HandleFunc("/musictop", music.FindMusicTop) //http://192.168.102.13:8080/musictop?ids=25fe4763-f363-4117-b4fb-fc01c6124cc4&userkey=huxn
	http.HandleFunc("/musiclrctop", music.MusicLrcHandler)
	http.ListenAndServe(":8080", nil)
}
