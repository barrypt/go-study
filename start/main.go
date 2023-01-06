package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	_ "time"
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

type Stu struct {
	A string
	B string
}
type StuGen[T any] struct {
	A T
	B string
}

func (stu *StuGen[T]) Get() {

	fmt.Println("A", stu.A)
}

func GenericGet[T int | string](aa T) {

	fmt.Println("TA", aa)

}

func AA(aa interface{}) (bb interface{}, err error) {

	switch res := aa.(type) {

	case *Stu:
		fmt.Println("stu", res)
	case string:
		fmt.Println("string", res)
	default:
		fmt.Println(aa)
	}

	return "", errors.New("111")

}

func insertCh(ch1 chan<- int) {

	for i := 0; i < 1; i++ {
		ch1 <- i
	}

	//close(ch1)
}
func readCh(ch1 <-chan int, cx context.Context) {

	for {

		select {

		case gg, ok := <-ch1:
			if !ok {
				fmt.Println("通道中已关闭")
				goto stop
			}
			fmt.Println("ch1", gg)
		case <-cx.Done():
			stu1 := Stu{A: "122", B: "3455"}
			fmt.Println("通道中已关闭", cx.Value(stu1))
			goto stop
		default:
			fmt.Println("通道中没有数据")
		}
		fmt.Println("waiting")
	}
stop:
}

type Str = int

var IsLoopback sync.WaitGroup

func main() {

	var c1 = make(chan int, 3)

	for i := 0; i < 3; i++ {

		c1 <- i

	}
	//close(c1)

	go func() {
		for g := range c1 {

			fmt.Println("g", g)

		}
	}()
	for y := 0; y < 10; y++ {

		IsLoopback.Add(1)
		go func(z int) {
			fmt.Println("zzz", z)
			IsLoopback.Done()
		}(y)

		go func() {
			IsLoopback.Wait()
			fmt.Println("yyy", y)
		}()
	}

	stu1 := Stu{A: "122", B: "3455"}
	stu2 := Stu{A: "122", B: "3455"}
	fmt.Println("stu1==stu2", stu1 == stu2)
	//xt := context.Background()
	//vacx := context.WithValue(xt, stu1, Str(236365))

	//chan1 := make(chan int)
	//sxt, cal := context.WithCancel(vacx)

	//time.AfterFunc(time.Duration(time.Second*1), func() {
	//	cal()
	//close(chan1)
	//})
	//go insertCh(chan1)
	//go readCh(chan1, sxt)

	gg := &StuGen[string]{A: "123456", B: "123"}

	fmt.Println("gg", gg)

	gg.Get()

	GenericGet(123)

	AA(&Stu{A: "122", B: "3455"})
	AA("234455")

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
