package main

import (

	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
    engine := gin.New()
	gin.SetMode(gin.DebugMode)
    engine.NoRoute(func(context *gin.Context) {
        requestPath := context.Request.URL.Path

        if strings.HasPrefix(requestPath, "/") {
            targetUrl, _ := url.Parse("http://localhost:8088")
            proxy := httputil.NewSingleHostReverseProxy(targetUrl)
            proxy.ServeHTTP(context.Writer, context.Request)
        } 

    })

	 engine.Run(":8082")
	
}