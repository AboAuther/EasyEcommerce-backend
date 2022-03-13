package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const DefaultPort = ":9999"

func main() {

	r := gin.Default()

	r.GET("/index/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
		c.JSON(http.StatusOK, gin.H{})
	})
	if err := r.Run(); err != nil {
		panic("run failed!")
	}
}
