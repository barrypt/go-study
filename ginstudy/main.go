package main

import (
	router "GO-STUDY/ginstudy/route"
	"encoding/json"
	"fmt"
	_ "net"
	"net/http"
	"os"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

type a struct {
	Name string `json:"name"`
}

func aa(a int) (t int) {

	t = a
	defer func() { t += 2 }()

	t += 10
	return t

}

func main() {

	fmt.Println("1", os.Getenv("AA"))
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

		if fileerr != nil {
			fmt.Println(fileerr)
		}
		ofile, errr := file.Open()
		if errr == nil {
			fmt.Println(errr)
		}
		defer ofile.Close()

		exfile, err := os.OpenFile("./aa.txt", int(os.O_CREATE), 0666)
		if err == nil {
			fmt.Println(err)
		}
		defer exfile.Close()
		wbuyte := make([]byte, 11)
		readf, errrr := ofile.Read(wbuyte)
		if errrr != nil {
			fmt.Println(errrr)
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
			fmt.Println(err)
		}
		fmt.Println(string(jsn))

	})
	r.GET("mumianhua", func(ctx *gin.Context) {

		Head := ctx.Request.Header
		tokne := Head.Get("token")
		data1 := `{"code":0,"msg":"成功","data":{"orgName":"测试渠道三","phone":"18513390121",
		"name":"哦婆婆","firstOrgId":3226735873409168,"firstOrgName":"北京分公司","orgId":3302629329677568},
		"timestamp":1670321150961,"success":true}`
		data2 := `{"code":0,"msg":"成功","data":{"orgName":"","phone":"18513391121",
		"name":"哦婆婆","firstOrgId":3226735873409168,"firstOrgName":"测试","orgId":3302629329677568},
		"timestamp":1670321150961,"success":true}`

		switch tokne {
		case "1":
			ff, err := simplejson.NewJson(([]byte(data1)))
			if err != nil {
				fmt.Println(err)
			}
			ctx.JSON(200, ff)
		case "2":
			ff1, err := simplejson.NewJson(([]byte(data2)))
			if err != nil {
				fmt.Println(err)
			}
			ctx.JSON(200, ff1)
		}

	})
	//gin.SetMode(gin.ReleaseMode)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
