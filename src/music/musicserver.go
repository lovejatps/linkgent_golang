// musicserver
package music

import (
	//"dataStruct"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"net/http"
	//	"strings"
	//"reflect"
)

const Url = "http://192.168.100.134:9200"

/**
*音乐查询接口1a
 */
func MusicHandler(w http.ResponseWriter, r *http.Request) {
	client, err := elastic.NewClient(elastic.SetURL(Url), elastic.SetMaxRetries(10))
	if err != nil {
		fmt.Fprint(w, "ES联接有问题", err)
	} else {
		jsdata := r.FormValue("jsondata")
		fmt.Println(jsdata)
		var conds Conditions
		err := json.Unmarshal([]byte(jsdata), &conds)
		if len(conds.UserKey) <= 0 {
			fmt.Fprint(w, "请出示你的Key")
		} else {
			fmt.Println("---->", conds.Special, "-->", len(conds.Special))
			if err != nil {
				fmt.Println("格式存在问题")
			} else {
				fmt.Fprint(w, getData(*client, conds))
			}
		}

	}

}

func getData(client elastic.Client, cond Conditions) string {

	searchService := client.Search()
	searchService = searchService.SearchType("dfs_query_then_fetch")
	searchService = searchService.Index("music")
	if len(cond.All) > 0 {
		searchService = searchService.Query(elastic.NewQueryStringQuery(cond.All))
	} else {
		if len(cond.ArtisName) > 0 {
			//	elastic.NewWildcardQuery("Name", "*hu*").Boost(1.2)
			searchService = searchService.Query(elastic.NewWildcardQuery("artisName", "*"+cond.ArtisName+"*"))
		}
		if len(cond.Name) > 0 {
			searchService = searchService.Query(elastic.NewWildcardQuery("name", "*"+cond.Name+"*"))
		}
		if len(cond.Special) > 0 {
			searchService = searchService.Query(elastic.NewWildcardQuery("special", "*"+cond.Special+"*"))
		}
	}
	serarchResult, err := searchService.From(0).Size(20).Pretty(true).Timeout("2s").Do()
	if err != nil {
		return "传入值有问题，没能找到你想要的歌曲"
	} else {
		items := make([]interface{}, 1, 1)
		var item interface{} = ""
		i := 0
		for _, hit := range serarchResult.Hits.Hits {
			if hit.Index != "music" {
				fmt.Errorf("expected SearchResult.Hits.Hit.Index = %q; got %q", "music", hit.Index)
			}
			err := json.Unmarshal(*hit.Source, &item)
			if err != nil {
				fmt.Println(err)
			}
			items = append(items, item)
			if i >= 19 {
				break
			}
			i = i + 1
		}
		for _, s := range items {
			fmt.Println(s)
		}
		lang, errjs := json.Marshal(items)
		if errjs != nil {
			fmt.Println("JSON", errjs)
		}
		return string(lang)

	}
	return "没能找到你需要的歌曲"
}

type Tweet struct {
	Name string "name"
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
	ArtisName string "artisName"
	//	ArtistGender string "artistGender"
}
type Conditions struct {
	ArtisName string "artisName"
	Special   string "special"
	Name      string "name"
	Lyric     string "lyric"
	All       string "all"
	UserKey   string "userkey"
}
