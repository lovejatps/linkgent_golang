// httpjson_test project main.go
package music

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/opesun/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	url_music       = "http://music.baidu.com/search/song?s=1&key="
	url_music_break = "&jump=0&start=0&size=10&third_type=0"
	url_music_top   = "http://ting.baidu.com/data/music/links?qq-pf-to=pcqq.temporaryc2c&songIds="
	//urllrc_top      = "http://music.baidu.com"
)

func GetMusicJosn(items []interface{}) string {
	var jsitems []interface{}
	var itemap map[string]interface{}
	for _, inte := range items {
		url := url_music_top + inte.(string)
		jsondata, err := simplejson.NewJson(fetch(&url))
		if err != nil {
			log.Fatalln(err)
		}
		itemap = make(map[string]interface{})

		jsondata = jsondata.Get("data").Get("songList").GetIndex(0)
		var item interface{} = ""
		itemap["albumName"] = jsondata.GetPath("albumName").MustString()
		itemap["artistGender"] = -1
		itemap["artistName"] = jsondata.GetPath("artistName").MustString()
		itemap["audioId"] = ""
		itemap["downloadNum"] = 0
		itemap["format"] = jsondata.GetPath("format").Interface()
		itemap["hotNum"] = 0
		itemap["lrcId"] = ""
		itemap["musiclrcurl"] = urllrc_top + jsondata.Get("lrcLink").MustString()
		itemap["musictopurl"] = jsondata.GetPath("songLink").MustString()
		itemap["name"] = jsondata.GetPath("songName").MustString()
		itemap["playNum"] = 0
		itemap["rate"] = jsondata.GetPath("rate").Interface()
		itemap["size"] = jsondata.GetPath("size").Interface()
		itemap["style"] = -1
		itemap["time"] = jsondata.GetPath("time").Interface()

		item = itemap
		jsitems = append(jsitems, item)

	}

	lang, errjs := json.Marshal(jsitems)
	if errjs != nil {
		fmt.Println(errjs)
	}
	//fmt.Println(string(lang))
	defer deletemap(itemap)
	return string(lang)
}

func deletemap(musicMap map[string]interface{}) {
	delete(musicMap, "albumName")
	delete(musicMap, "artistGender")
	delete(musicMap, "artistName")
	delete(musicMap, "audioId")
	delete(musicMap, "downloadNum")
	delete(musicMap, "downloadNum")
	delete(musicMap, "hotNum")
	delete(musicMap, "hotNum")
	delete(musicMap, "musiclrcurl")
	delete(musicMap, "musictopurl")
	delete(musicMap, "name")
	delete(musicMap, "playNum")
	delete(musicMap, "rate")
	delete(musicMap, "size")
	delete(musicMap, "style")
	delete(musicMap, "time")
}

func fetch(url *string) (html []byte) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		//html = string(robots)
		return robots
	} else {
		//	html = ""
	}
	return nil
}

func GetMusicID(name string) []interface{} {
	p, err := goquery.ParseUrl(url_music + name + url_music_break)
	if err != nil {
		fmt.Println(err)
	} else {
		var items []interface{}
		t := p.Find(".song-item-hook")
		if t.Length() > 0 {
			for i := 0; i < t.Length(); i++ {
				if strings.Contains(t.Eq(i).Html(), "data-songdata") {
					str := strings.SplitN(t.Eq(i).Html(), "data-songdata", 2)[0]
					str = strings.SplitN(str, "href=", 2)[1]
					if strings.Contains(str, " ") {
						str = strings.Split(str, " ")[0]
					} else {
						continue
					}
					if strings.Contains(str, "/") {
						item_ht := strings.Split(str, "/")
						str = item_ht[len(item_ht)-1]
					} else {
						continue
					}
					if strings.Contains(str, "\"") {
						str = strings.Split(str, "\"")[0]
					} else {
						continue
					}
					items = append(items, str)
				}
			}
		}
		return items

	}
	return nil
}
