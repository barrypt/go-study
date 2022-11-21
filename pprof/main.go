package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len:%d", Add("gwyysdfsdfsdfsdfsdfsdfsdfsd"))
			time.Sleep(time.Microsecond + 10)
		}
	}()
	//http://localhost:6060/debug/pprof/
	_ = http.ListenAndServe(":6060", nil)
}
func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}
