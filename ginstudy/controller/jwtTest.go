package controller

import (
	"GO-STUDY/ginstudy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	if !(username == "admin" && password == "123456") {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	user := utils.Users{
		Username: "admin",
		Password: "123456",
	}
	//生成token
	token, err := utils.GenToken(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{"token": token},
	})

}

func UserListController(ctx *gin.Context) {
	userClaims, _ := ctx.Get("claims")
	user := userClaims.(*utils.CustomClaims)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": user.Username,
	})
}
