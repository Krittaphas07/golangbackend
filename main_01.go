package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//------------GET-------------

	r.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "P nut ya ooor",
		})
	})

	//-----------POSE--------------

	r.POST("/topten", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "dadaiaijdajdjadjagtoperro",
		})
	})
	//----------PUT-------------------
	r.PUT("/MM", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "P nut ya ooor",
		})
	})

	//-----------DELETE--------------

	r.DELETE("/NUT", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "dadaiaijdajdjadjagtoperro",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}

