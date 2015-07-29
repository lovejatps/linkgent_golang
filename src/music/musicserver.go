// musicserver
package music

import (
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"net/http"
)

const urles = "http://192.168.100.134:9200"

var size = 20

/**
*音乐查询接口1a
 */
func MusicHandler(w http.ResponseWriter, r *http.Request) {
	client, err := elastic.NewClient(elastic.SetURL(urles), elastic.SetMaxRetries(10))
	if err != nil {
		fmt.Fprint(w, "ES联接有问题", err)
	} else {
		jsdata := r.FormValue("jsondata")

		var conds Conditions
		err := json.Unmarshal([]byte(jsdata), &conds)
		if err != nil {
			fmt.Println("格式存在问题")
		} else if len(conds.UserKey) <= 0 {
			fmt.Fprint(w, "请出示你的Key")

		} else {
			fmt.Println(conds.ArtisName, conds.Special, conds.Name)
			fmt.Fprint(w, getData(*client, conds))
		}
	}

}

func elResult(client elastic.Client, cond Conditions) (*elastic.SearchResult, error) {
	if len(cond.All) > 0 {
		return client.Search().
			Index("music2").
			SearchType("dfs_query_then_fetch").
			Query(elastic.NewQueryStringQuery(cond.All)).
			From(0).
			Size(size).
			Timeout("3s").
			Do()
	} else {
		qbool := elastic.NewBoolQuery()
		if len(cond.ArtisName) > 0 {
			q_artisName := elastic.NewQueryStringQuery(cond.ArtisName).AnalyzeWildcard(false).DefaultOperator("and")
			q_artisName = q_artisName.DefaultField("artistName")
			qbool = qbool.Must(q_artisName)
		}
		if len(cond.Name) > 0 {
			q_name := elastic.NewQueryStringQuery(cond.Name).AnalyzeWildcard(false).DefaultOperator("and")
			q_name = q_name.DefaultField("name")
			qbool = qbool.Must(q_name)
		}
		if len(cond.Special) > 0 {
			q_special := elastic.NewQueryStringQuery(cond.Special).AnalyzeWildcard(false).DefaultOperator("and")
			q_special = q_special.DefaultField("special")
			qbool = qbool.Must(q_special)
		}
		data, errdata := json.Marshal(qbool.Source())
		if errdata == nil {
			fmt.Println(string(data))
		}
		return client.Search().Index("music2").Query(qbool).From(0).Size(size).Timeout("3s").Do()
	}
	return nil, nil
}

func getData(client elastic.Client, cond Conditions) string {

	serarchResult, err := elResult(client, cond)

	if err != nil {
		return "传入值有问题，没能找到你想要的歌曲"
	} else {
		//items := make([]interface{}, 1, 1)
		var items []interface{}
		var item interface{} = ""
		for i, hit := range serarchResult.Hits.Hits {
			if hit.Index != "music" {
				fmt.Errorf("expected SearchResult.Hits.Hit.Index = %q; got %q", "music", hit.Index)
			}
			err := json.Unmarshal(*hit.Source, &item)
			if err != nil {
				fmt.Println(err)
			}
			itemap := item.(map[string]interface{})
			itemap["musictopurl"] = ""
			itemap["musiclrcurl"] = ""
			item = itemap
			items = append(items, item)
			items[i] = item
			fmt.Println("本次查询的条数是：", i)
			if i >= size-1 {
				break
			}
			i = i + 1
		}
		if len(items) <= 0 {
			fmt.Println("本地未能找到你需要的歌曲")
			var condition string
			if len(cond.ArtisName) > 0 {
				condition = cond.ArtisName
			}
			if len(cond.Name) > 0 {
				condition = condition + "+" + cond.Name
			}
			if len(cond.Special) > 0 {
				condition = condition + "+" + cond.Special
			}
			str := GetMusicJosn(GetMusicID(condition))
			return str
		} else {
			lang, errjs := json.Marshal(items)
			if errjs != nil {
				fmt.Println("JSON", errjs)
			}
			return string(lang)
		}

	}
	return "没能找到你需要的歌曲"
}

//type Tweet struct {
//Name string "name"
//	Style        string "style"
//	Ablum        string "ablum"
//	PublishDate  string "publishDate"
//	Rate         string "rate"
//	size         string "size"
//	AudioId      string "audioID"
//	LrcID        string "lrcID"
//	DownloadNum  string "downloadNum"
//	PlayNum      string "playNum"
//	HotNum       string "hotNum"
//ArtisName string "artisName"
//	ArtistGender string "artistGender"
//}
type Conditions struct {
	ArtisName string "artisName"
	Special   string "special"
	Name      string "name"
	Lyric     string "lyric"
	All       string "all"
	UserKey   string "userkey"
}
