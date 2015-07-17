// findMusicTop
package music

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"strings"
)

const (
	url_mg  = "192.168.102.194,192.168.102.195,192.168.102.196"
	mongodb = "Music"
)

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func FindMusicTop(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	parameter := make(map[string]string)
	for k, v := range r.Form {
		parameter[k] = strings.Join(v, "")
	}
	session, err := mgo.Dial(url_mg)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(mongodb)

	Id := parameter["ids"]
	Userkey := parameter["userkey"]
	Ids := strings.Split(parameter["ids"], ",")
	//读取二进制文件
	if strings.EqualFold(Id, "") {
		fmt.Fprintf(w, "歌曲不能为空")
	} else {
		if strings.EqualFold(Userkey, "") {
			fmt.Fprintf(w, "您的key值不合法")
		} else {
			for i := 0; i < len(Ids); i++ {
				file, err := db.GridFS("fs").OpenId(Ids[i])
				//Open("2.mp4")
				if err == nil {
					b := make([]byte, file.Size())
					check(err)
					m, err := file.Read(b)
					fmt.Print("m=", m)
					check(err)
					_, err = w.Write(b)
				} else {
					fmt.Fprintf(w, "没有查询到你需要的资源")
				}
			}
		}
	}
}
