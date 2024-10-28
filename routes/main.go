package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run(port string) {
	v1 := router.Group("api/v1")
	addAuthRoute(v1)
	router.Run(port)
}
