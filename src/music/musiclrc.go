// musiclrc
package music

import (
	//	"encode"
	"errors"
	"fmt"
	"github.com/opesun/goquery"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const url_lrc = "http://music.baidu.com/search/lrc?key="
const urllrc_top = "http://music.baidu.com"

func MusicLrcHandler(w http.ResponseWriter, r *http.Request) {
	userkey := r.FormValue("userkey")
	if len(userkey) > 0 {
		lrc_name := r.FormValue("song")
		if len(lrc_name) > 0 {
			fmt.Println(r.FormValue("song"))
			site, err := getHmtl(url_lrc + lrc_name)
			if err == nil {
				if len(site) > 0 {
					client := &http.Client{}
					req, err := http.NewRequest("GET", urllrc_top+site, nil)
					if err != nil {
						fmt.Fprint(w, "没能找到你需要的歌词")
					}
					resp, err := client.Do(req)
					if err != nil {
						fmt.Fprint(w, "没能找到你需要的歌词")
					}
					if resp.StatusCode == 200 {
						lcrbayd, err := ioutil.ReadAll(resp.Body)
						resp.Body.Close()
						if err != nil {
							fmt.Fprint(w, "没能找到你需要的歌词")
						}
						w.Header().Set("size", strconv.Itoa(len(lcrbayd)))
						w.Write(lcrbayd)

					}
				}
			}

		}
	} else {
		fmt.Fprint(w, "非法用户")
	}
}

func getHmtl(url string) (string, error) {
	p, err := goquery.ParseUrl(url)
	if err != nil {
		return "", err
	} else {
		t := p.Find(".lyric-action")
		if t.Length() > 0 {
			for i := 0; i < 1; i++ { //只返回一条记录
				if strings.Contains(t.Html(), "down-lrc-btn") {
					lrc := strings.Split(strings.SplitN(t.Html(), "down-lrc-btn", 2)[1], "&#39;")
					return lrc[3], nil

				}

			}
		}

	}
	return "", errors.New("没能找到你需要的歌词")
}
