// RadioHandler
package main

import (
	"fmt"
	//"github.com/olivere/elastic"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")
}

const (
	urles    = "http://192.168.100.134:9200"
	indexing = "audio"
)

type RadioTpy struct {
	Cid  string "cid"
	Page string "page"
}

/**
*处理电台分类
 */
func RadioHandler(w http.ResponseWriter, r *http.Request) {
	//	client, err := elastic.NewClient(elastic.SetURL(urles), elastic.SetMaxRetries(10))
	//	if err != nil {
	//		fmt.Fprint(w, "ES连接有问题")
	//	}

	//	cid := r.FormValue("cid")
	//	if nil == cid {
	//		fmt.Fprint(w, "格式有问题")
	//	} else if "-1" == cid {

	//	}

}
