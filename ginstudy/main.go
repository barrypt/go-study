package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"GO-STUDY/ginstudy/route"
	"github.com/gin-gonic/gin"

)

type a struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	router.Router(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("test", func(ctx *gin.Context) {
		a := ctx.Query("a")
		b := ctx.Query("b")
		aint, _ := strconv.Atoi(a)
		bint, _ := strconv.Atoi(b)
		fmt.Println(aint + bint)
		ctx.JSON(http.StatusOK, aint+bint)
	})
	r.POST("postData", func(ctx *gin.Context) {
		a := &a{}
		file, fileerr := ctx.FormFile("file")

		if fileerr == nil {

		}
		ofile, errr := file.Open()
		defer ofile.Close()

		if errr == nil {

		}
		exfile, err := os.OpenFile("./aa.txt", int(os.O_CREATE), 0666)
		defer exfile.Close()
		wbuyte := make([]byte, 11)
		readf, errrr := ofile.Read(wbuyte)
		if errrr == nil {

		}
		fmt.Println("wbuyte", string(wbuyte))
		fmt.Println("readf", readf)
		exfile.Write(wbuyte)
		name, boo := ctx.GetPostForm("name")
		if boo {
			fmt.Println(name)
		}
		jsn, err := json.Marshal(a)
		if err == nil {

		}
		fmt.Println(string(jsn))

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
