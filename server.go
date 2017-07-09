package main

import(
	//"github.com/nobonobo/unqlitego"
	// "github.com/tpotlog/unqlitego"
	// "github.com/tpotlog/unqlitego/collections"
	//"github.com/nobonobo/unqlitego/collections"
	//"collections"
	"github.com/gin-gonic/gin"
	// "fmt"
	//"errors"
)
func ping (c *gin.Context) {
   c.JSON(200, gin.H{
	   "message": "pong",
   })
}
func pipu (c *gin.Context) {
   c.JSON(200, gin.H{
	   "pipu": "pipu bueno",
   })
}

func main() {
	r := gin.Default()
	r.GET("/ping",ping)
	r.GET("/pipu",pipu)
	r.Run() // listen and server on 0.0.0.0:8080
}
