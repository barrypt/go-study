package main

import (
	"fmt"
	"runtime"
	"time"
)

func GetData() {
	var ch chan struct{}
	go func() {
		<-ch
	}()
}

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	GetData()
	time.Sleep(2 * time.Second)
}
