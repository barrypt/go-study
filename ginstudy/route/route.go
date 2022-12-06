package router

import (
   "GO-STUDY/ginstudy/controller"
   _ "GO-STUDY/ginstudy/middleware"
   "github.com/gin-gonic/gin"
)

func Router(r *gin.Engine)  {

   //用户登录
   r.GET("/login", controller.LoginController)
   //使用中间件
   //r.Use(middleware.JWTAuth())
   //获取列表数据
   r.GET("/list", controller.UserListController)
}
