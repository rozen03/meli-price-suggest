package main

import "github.com/gin-gonic/gin"

//"github.com/nobonobo/unqlitego"
// "github.com/tpotlog/unqlitego"
// "github.com/tpotlog/unqlitego/collections"
//"github.com/nobonobo/unqlitego/collections"
//"collections"
// "fmt"
//"errors"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func pipu(c *gin.Context) {
	c.JSON(200, gin.H{
		"pipu": "pipu bueno",
	})
}

func mainn() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/pipu", pipu)
	r.Run() // listen and server on 0.0.0.0:8080
}
