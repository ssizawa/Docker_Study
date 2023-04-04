package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/hoka-isdl/docker_handson/server/controller"
	
)

func main() {

	//Ginの変数
	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	// //routerを渡す
	controller.Router(router)

	port := os.Getenv("CONTAINER_SERVER_PORTS")
	//ポートを指定して実行
	router.Run(":"+ port)
}
