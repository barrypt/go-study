package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
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
