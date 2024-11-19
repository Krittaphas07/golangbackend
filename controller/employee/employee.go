package employee

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

//GET
func GET(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}

//POST
func POST(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}

//PUT
func PUT(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}

//DELETE
func DELETE(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}
