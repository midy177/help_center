package biz

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

func UploadImage(reader io.Reader, fn string) *BaseJson {
	url := strconv.FormatInt(time.Now().UnixMilli(), 10) + path.Ext(fn)
	fileName := path.Join("static/upload", url)
	f, err := os.Create(fileName)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		var b []byte
		b, err = ioutil.ReadAll(reader)
		if err != nil {
			log.Println(err.Error())
			return &BaseJson{Code: 0, Data: err.Error()}
		}
		_, err = f.Write(b)
		if err != nil {
			log.Println(err.Error())
			return &BaseJson{Code: 0, Data: err.Error()}
		}
		return &BaseJson{Code: 1, Data: "/upload/" + url}
	}
}
