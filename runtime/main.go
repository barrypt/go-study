package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/arl/statsviz"
)

func main() {

	fmt.Println("toutine", runtime.NumGoroutine())
	fmt.Println("toutine", runtime.NumCPU())
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	memstr, _ := json.Marshal(mem)
	fmt.Println("mem", string(memstr))

	//http://localhost:6060/debug/statsviz/
	statsviz.RegisterDefault()
	log.Println(http.ListenAndServe(":6060", nil))

}
