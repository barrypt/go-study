package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/stringx"
)

func main() {

	filter := stringx.NewTrie([]string{
		"AV演员",
		"苍井空",
		"AV",
		"日本AV女优",
		"AV演员色情",
	}, stringx.WithMask('?'))
	safe, keywords, found := filter.Filter("日本AV演员兼电视、电影演员。苍井空AV女优是xx出道, 日本AV女优们最精彩的表演是AV演员色情表演")
	fmt.Println(safe)
	fmt.Println(keywords)
	fmt.Println(found)


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
