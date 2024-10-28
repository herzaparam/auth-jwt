package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAuthRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, "login")
	})
	auth.GET("/refresh-token", func(c *gin.Context) {
		c.JSON(http.StatusOK, "refresh token")
	})
}
