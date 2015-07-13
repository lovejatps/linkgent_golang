// weather
package weather

import (
	"citytool"
	"dataStruct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//const Key = "IEG1K5BBTK"
const Key = "FSUZWU8RKI"

var weatherMap map[string]dataStruct.Data

/**
*天气接口
 */
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	var weather dataStruct.Data
	var getOK bool
	parameter := Parameter{
		//cityName :encode.Decode(r.FormValue("city")),
		cityName: r.FormValue("city"),
		cityId:   citytool.GetCityID(r.FormValue("city")),
	}
	if strings.EqualFold(parameter.cityId, "null") {
		fmt.Fprint(w, "传入数据有问题")
	}

	weather, getOK = sendWeather(parameter)
	if getOK {
		jsondata, err := json.Marshal(weather)
		if err != nil {
			fmt.Fprint(w, err)
		} else {
			fmt.Fprint(w, string(jsondata))
		}
	} else {
		fmt.Fprint(w, "传入数据有问题")
	}
}

type Parameter struct {
	cityName string
	cityId   string
}

func sendWeather(p Parameter) (weather dataStruct.Data, ok bool) {
	if !strings.EqualFold(p.cityId, "null") {

		if weatherMap == nil {
			weatherMap = make(map[string]dataStruct.Data)
			weather := getWeatherData(p)
			weatherMap[p.cityId] = weather
			return weather, true
		} else {
			history, ok := weatherMap[p.cityId]
			if ok {
				if isTime(history.Weather[0].Last_update) {
					return history, true
				} else {
					weather := getWeatherData(p)
					weatherMap[p.cityId] = weather
					return weather, true
				}
			} else {
				weather := getWeatherData(p)
				weatherMap[p.cityId] = weather
				return weather, true
			}
		}

	} else {
		return weather, false
	}

	return weather, false

}

func getWeatherData(p Parameter) dataStruct.Data {
	var weather dataStruct.Data
	url := "https://api.thinkpage.cn/v2/weather/all.json?city="
	url = url + p.cityId
	url = url + "&language=zh-chs&unit=c&aqi=city&key=" + Key
	fmt.Println(url)
	response, _ := http.Get(url)
	defer response.Body.Close()
	jsonbody, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(jsonbody, &weather)
	if err != nil {
		fmt.Println("error:", err)
	}
	return weather

}

func isTime(s string) bool {
	str1 := strings.Split(s, ":")
	str1 = strings.Split(str1[0], "T")
	date1 := str1[0]
	t := time.Now().Unix()
	timestr := time.Unix(t, 0).String()
	date2 := strings.Fields(timestr)[0]
	if strings.EqualFold(date1, date2) {
		time_h, err := strconv.Atoi(str1[1])
		if err != nil {
			fmt.Println(err)
			return false
		} else {
			current_hour, err1 := strconv.Atoi(strings.Split(strings.Fields(timestr)[1], ":")[0])
			if err1 != nil {
				fmt.Println(err1)
				return false
			} else {
				if current_hour == time_h {
					return true
				} else {
					return false
				}
			}

		}

	} else {
		return false
	}

	return false
}
