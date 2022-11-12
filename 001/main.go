package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	m      = iota
	w      = "woman"
	unknow = iota
)
const (
	MONDAY = iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

func insertCh(ch1 chan int) {

	for i := 0; i < 3; i++ {
		ch1 <- i
	}

	close(ch1)
}
func readCh(ch1 chan int) {

	for {

		select {

		case gg, ok := <-ch1:
			if !ok {
				fmt.Println("通道中已关闭")
				goto stop
			}
			fmt.Println("ch1", gg)
		default:
			fmt.Println("通道中没有数据")
		}
		fmt.Println("waiting")
	}
stop:
}
func main() {

	today := time.Date(2022, 10, 31, 0, 0, 0, 0, time.Local)
	nextDay := today.AddDate(0, 1, 0)
	fmt.Println(nextDay.Format("20060102"))
	// 输出：20221201

	var ch1 = make(chan int, 2)

	go insertCh(ch1)

	go readCh(ch1)

	var ary = [3]int{1}

	for i := 0; i < len(ary); i++ {
		ary[i] = i
	}

	for i, v := range ary {

		fmt.Println("sdsd", i, v)
	}

	var mm map[string]int = make(map[string]int)
	mm["1"] = 1
	hh := mm["1"]
	fmt.Println("hh:", hh)
	fmt.Println("mmlen:", len(mm))
	//fmt.Println(cap(mm))

	for k, v := range mm {

		fmt.Println("kv", k, v)

	}

	fmt.Println(m)
	fmt.Println(w)
	fmt.Println(unknow)
	fmt.Println(TUESDAY)

	fmt.Println("SEED", time.Now().UnixMilli())
	rand.Seed(time.Now().Unix())
	rd := rand.Intn(100)
	fmt.Println("rand", rd)

	var i interface{} = "hello"

	switch i.(type) {
	case string:
		fmt.Println("string:", i)
	case float64:
		fmt.Println("float:", i)
	}

	f, ok := i.(float64) //  no runtime panic
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	//fmt.Println(f)

	http.Handle("/", http.FileServer(getFileSystem(false)))
	ip, err := getLocalIP()
	if err != nil {
		return
	}
	log.Println("启动成功，通过 http://" + ip + ":10002 访问")

	server := http.Server{
		Addr:    ":10002",
		Handler: nil,
	}
	server.ListenAndServe()
}

//go:embed wwwroot
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		return http.FS(os.DirFS("wwwroot"))
	}

	fsys, err := fs.Sub(embededFiles, "wwwroot")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func getLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}
