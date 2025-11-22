package routers

import (
	"gin-todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 它的意思是：只有访问 "/" (根路径) 时，才去读取 "./static/index.html"
	// 这不是通配符，所以不会和 /todo 冲突
	r.StaticFile("/", "../static/index.html")
	// 注册路由
	r.GET("/todo", controllers.GetTodos)
	r.POST("/todo", controllers.CreateTodo)
	r.PUT("/todo/:id", controllers.UpdateTodo)
	r.DELETE("/todo/:id", controllers.DeleteTodo)

	return r
}
