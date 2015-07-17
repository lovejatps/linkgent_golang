package dataStruct

import ()

type Data struct {
	Status  string
	Weather []Weatherdata
}

type Weatherdata struct {
	City_name   string
	City_id     string
	Last_update string
	Now         NowWeather
	Today       Todaydata
	Future      []Futuredata
}

type NowWeather struct {
	Text            string
	Code            string
	Temperature     string
	Feels_like      string
	Wind_direction  string
	Wind_speed      string
	Wind_scale      string
	Humidity        string
	Visibility      string
	Pressure        string
	Pressure_rising string
	Air_quality     string
}

type Todaydata struct {
	Sunrise    string
	Sunset     string
	Suggestion string
}

type Futuredata struct {
	Date  string
	Day   string
	Text  string
	Code1 string
	Code2 string
	High  string
	Low   string
	Cop   string
	Wind  string
}

type MusicData struct {
	ID        string //
	CdTitle   string //专辑名
	Tory      string //发行年份
	SongID    string //歌曲ID
	Song      string //歌曲名
	Character string //品质
	Size      string //文件大小
	LrcID     string //歌词ID
	Lrc       string //歌词
	Edloaded  string //歌曲下载量
	Spreads   string //歌曲播放量
	Artist    string //歌手名
	Gender    string //性别
	Hot       string //歌手热度
}
