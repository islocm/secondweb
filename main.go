package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var router *gin.Engine

func main() {

	e := connect()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("ist work")
	router = gin.Default()
	router.GET("/", index)
	router.Static("/assets", cfg.Assets)
	router.LoadHTMLFiles(cfg.HTML + "index.html")

	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
