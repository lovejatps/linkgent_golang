package main

import (
	"music"
	"net/http"
	"weather"
)

func main() {
	http.HandleFunc("/weather", weather.DefaultHandler)
	http.HandleFunc("/music", music.MusicHandler)
	http.HandleFunc("/musictop", music.FindMusicTop) //?ids=xxx&userkey=xxx   http://192.168.102.13:8080/musictop?ids=25fe4763-f363-4117-b4fb-fc01c6124cc4&userkey=huxn
	http.HandleFunc("/musiclrctop", music.MusicLrcHandler)
	http.ListenAndServe(":8080", nil)
}
