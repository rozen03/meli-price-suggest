package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func prices(c *gin.Context, ch chan Cosa) {
	id := c.Param("id")
	res := Suggest(id, func(category string) obtainedData { return PreciosYVentas(category, ch) })
	c.JSON(200, gin.H{
		"max":       res.max,
		"suggested": res.suggested,
		"min":       res.min,
	})
}
func start() {
	rand.Seed(time.Now().Unix())
	ch := make(chan Cosa)
	go scheduler(ch)
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/categories/:id/prices", func(c *gin.Context) { prices(c, ch) })
	r.Run(":8081")

}
