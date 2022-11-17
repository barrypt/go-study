package middleware

import (
   "errors"
   "GO-STUDY/ginstudy/utils"
   "github.com/gin-gonic/gin"
   "log"
   "net/http"
   "strings"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
   return func(ctx *gin.Context) {
   	authHeader := ctx.Request.Header.Get("Authorization")
   	if authHeader == "" {
   		ctx.JSON(http.StatusOK, gin.H{
   			"code": -1,
   			"msg": "无权限访问，请求未携带token",
   		})
   		ctx.Abort()//结束后续操作
   		return
   	}
   	log.Print("token:", authHeader)

   	//按空格拆分
   	parts := strings.SplitN(authHeader, " ", 2)
   	if !(len(parts) == 2 && parts[0] == "Bearer")  {
   		ctx.JSON(http.StatusOK, gin.H{
   			"code": -1,
   			"msg":  "请求头中auth格式有误",
   		})
   		ctx.Abort()
   		return
   	}

   	//解析token包含的信息
   	claims ,err := utils.ParseToken(parts[1])
   	if err != nil {
   		ctx.JSON(http.StatusOK, gin.H{
   			"code": -1,
   			"msg":  "无效的Token",
   		})
   		ctx.Abort()
   		return
   	}

   /*	if err := CheckUserInfo(claims); err != nil{
   		ctx.JSON(http.StatusOK, gin.H{
   			"code": -1,
   			"msg":  "用户名或密码错误",
   		})
   		ctx.Abort()
   		return
   	}*/

   	// 将当前请求的claims信息保存到请求的上下文c上
   	ctx.Set("claims", claims)
   	ctx.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息

   }
}

//检查用户名信息
func CheckUserInfo(claims *utils.CustomClaims) error {
   username := claims.Username
   password := claims.Password
   //获取数据库用户名及密码
   if  username == "admin" && password == "123456"{
   	return nil
   }
   return errors.New("用户名或密码错误")
}
