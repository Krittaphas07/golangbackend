package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET Method!",
	})
}

// POST
func POST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee POST Method!",
	})
}

// PUT
func PUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee PUT Method!",
	})
}

// DELETE
func DELETE(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee DELETE Method!",
	})
}

// GET By ID
func GETEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

// POST By ID
func POSTEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

// PUT By ID
func PUTEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

// DELETE By ID
func DELETEEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}
