package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//Employee API method
	r.GET("/employee", getEmployee)
	r.POST("/topten", postEmployee)
	r.PUT("/MM", putEmployee)
	r.DELETE("/NUT", deleteEmployee)


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}

func getEmployee(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
  "message": "Employee GET Method!",
  })
}

func postEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "Topten GET Method!",
	})
  }

func putEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "MM GET Method!",
	})
  }

func deleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "NUT GET Method!",
	})
  }