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
			"message": "Get Employee",
		})
	})

	//-----------POSE--------------

	r.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Past employee",
		})
	})
	//----------PUT-------------------
	r.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Put Employee",
		})
	})

	//-----------DELETE--------------

	r.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete Employee",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
