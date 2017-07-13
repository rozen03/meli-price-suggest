package main

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func prices(c *gin.Context) {
	id := c.Param("id")
	res := Suggest(id, PreciosYVentas)
	c.JSON(200, gin.H{
		"max":       res.max,
		"suggested": res.suggested,
		"min":       res.min,
	})
	// c.String(http.StatusOK, "Hello %s", id)
}
func start() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/categories/:id/prices", prices)
	r.Run(":8081") // listen and server on 0.0.0.0:8081
}
