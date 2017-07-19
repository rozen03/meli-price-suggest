package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const melink = "https://api.mercadolibre.com/sites/MLA/search?limit=200&category="

// func ping(c *gin.Context) {
// c.JSON(200, gin.H{
// "message": "pong",
// })
// }
func prices(c *gin.Context, ch chan ArgsAndResult) {
	res := Suggest(c.Param("id"), ch, GetMeli)
	c.JSON(200, gin.H{
		"max":       strconv.FormatFloat(res.max, 'f', 2, 64),
		"suggested": strconv.FormatFloat(res.suggested, 'f', 2, 64),
		"min":       strconv.FormatFloat(res.min, 'f', 2, 64),
	})
}
func GetMeli(args string) (*http.Response, error) { return http.Get(melink + args) }

func start() {
	ch := startWorkers(maxChanelsSched)
	r := gin.Default()
	// r.GET("/ping", ping)
	r.GET("/categories/:id/prices", func(c *gin.Context) { prices(c, ch) })
	r.Run(":8081")

}
