package routes

import (
	"github.com/challenge-bw-go/src/controllers"
	"github.com/gin-gonic/gin"
)

func HelloRoutes(route *gin.RouterGroup) {
	group := route.Group("/hello")
	group.GET("/", controllers.HelloWorld)
}
