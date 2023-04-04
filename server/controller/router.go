package controller

import (
	//HTTPクライアントとサーバーの実装
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoka-isdl/docker_handson/server/repository"
)

func Router(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", gin.H{
			"judge": true,
		})
	})

	router.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		repository.Register(username,password)
		
		c.HTML(http.StatusOK, "index.html", gin.H{
			"judge": true,
			"username": username,
		})
	})

}