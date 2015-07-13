package encode

import (
	"bytes"
    "text-master/encoding/simplifiedchinese"
    "text-master/transform"
    "io/ioutil"
)

func Encode(src string) (dst string) {
    data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
    if err == nil {
        dst = string(data)
    }
    return
}
func Decode(src string) (dst string) {
    data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
    if err == nil {
        dst = string(data)
    }
    return
}