package citytool

import (
	"configs"
	//"fmt"
	"strings"
)

var citys string
var strarray = [...]string{
	"傈傈族",
	"彝族",
	"藏族",
	"土族",
	"布依族",
	"苗族",
	"白族",
	"满族",
	"瑶族",
	"黎族",
	"侗族",
	"土家族",
	"朝鲜族",
	"壮族",
	"回族",
	"蒙古族",
	"傣族",
	"佤族",
	"拉祜族",
	"布朗族",
	"景颇族",
}

func GetCityID(city string) string {
	ciarray := [...]string{"市", "县", "自治县", "区", "自治区", "自治州"}
	cityname := cityname(city)
	cityname = srtreplace(cityname)
	cityID, ok := configs.CityMap[city]
	if ok {
		return cityID
	} else {
		for _, value := range ciarray {
			//	fmt.Println("cityN", value)
			if strings.HasSuffix(cityname, value) {
				cityID, ok = configs.CityMap[strings.SplitN(cityname, value, 2)[0]]
				if ok {
					return cityID
				} else {
					cityID, ok = configs.CityMap[cityname]
					if ok {
						return cityID
					}
				}
			}

		}
	}
	return "没找到你提供的城市"
}

//把所有特定的词组去掉
func srtreplace(srt string) string {
	racestr := srt
	for _, value := range strarray {
		racestr = strings.Replace(racestr, value, "", -1)
	}
	//	fmt.Println("-----------------", racestr)
	return racestr
}

func cityname(city string) string {
	if strings.HasSuffix(city, "县") { //以“县”结尾
		if strings.Contains(city, "市") {
			citys = strings.SplitAfterN(city, "市", 2)[1]
		}
	}
	if strings.HasSuffix(city, "市") {
		if strings.Count(city, "市") >= 2 {
			citys = strings.SplitAfterN(city, "市", 2)[1]
		} else {
			if strings.Contains(city, "省") {

				citys = strings.SplitAfterN(city, "省", 2)[1]
			}
		}
	}
	if strings.Contains(city, "区") {
		if strings.HasSuffix(city, "区") {
			if strings.Count(city, "区") >= 2 {
				citys = strings.SplitAfterN(city, "区", 2)[1]
			} else {
				if strings.Contains(city, "特别行政区") {
					citys, _ = SubstringIndex(city, "特别行政区")

				}
				if strings.Contains(city, "市") && !strings.HasSuffix(city, "市") {
					citys = strings.SplitAfterN(city, "市", 2)[1]
				}
				if !strings.Contains(city, "市") {
					citys = city
				}
			}
		} else {
			citys, _ = SubstringEnd(city, "区")
		}
	}
	if strings.Contains(city, "自治州") {
		if strings.HasSuffix(city, "自治州") && strings.Contains(city, "省") {
			citys = strings.SplitAfterN(city, "省", 2)[1]
		} else if strings.HasSuffix(city, "自治州") && !strings.Contains(city, "省") {
			citys = city
		} else {
			citys, _ = SubstringEnd(city, "自治州")
		}

	}

	return citys
}

func getcityname(city string) string {

	if strings.Contains(city, "县") {
		citys = strings.SplitAfterN(city, "市", 2)[1]
	}
	if strings.Contains(city, "市") {

		if strings.Count(city, "市") >= 2 {
			citys = strings.SplitAfterN(city, "市", 2)[1]
		} else {
			citys = strings.SplitAfterN(city, "省", 2)[1]
		}

	}
	if strings.Contains(city, "区") {
		if strings.Count(city, "区") >= 2 {
			citys = strings.SplitAfterN(city, "区", 2)[1]
		} else {
			if strings.Contains(city, "特别行政区") {
				citys, _ = SubstringIndex(city, "特别行政区")

			}
			if strings.Contains(city, "市") {
				citys = strings.SplitAfterN(city, "市", 2)[1]
			}
		}

	}
	return ""
}

//从字符出现位置开始
func SubstringEnd(str, substr string) (string, bool) {
	i := UnicodeIndex(str, substr)
	rs := []rune(str)
	lth := len(rs)
	sublen := len([]rune(substr))
	if i != -1 && i < lth {
		return string(rs[i+sublen : lth]), true
	}
	return "传入的参数有问题", false

}

//从头开始
func SubstringIndex(str, substr string) (string, bool) {
	i := UnicodeIndex(str, substr)
	rs := []rune(str)
	lth := len(rs)
	if i != -1 && i < lth {
		return string(rs[0:i]), true
	}
	return "传入的参数有问题", false
}

//计算字符出现位置
func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}
