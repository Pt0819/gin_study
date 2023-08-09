package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index gin框架提供上下文操作，以下示例中的context在正式开发中会返回很多信息
func Index(context *gin.Context) {
	context.String(200, "hello_world")
}

func main() {
	//创建默认路由
	router := gin.Default()

	//绑定路由规则和路由函数，访问/index的路由，交给对应的函数处理
	//(group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes
	router.GET("/index", Index)

	//启动监听，gin会把web服务运行在本机的0.0.0.0:8080上
	router.Run(":8080")

	//router.Run()本质上就是http.ListenAndServe的进一步封装
	http.ListenAndServe(":8080", router)
}
