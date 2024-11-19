package main

import (
  EmployeeController "backend/api/controller/employee"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  //GET
  r.GET("/employee", EmployeeController.GET)
  //POST
  r.POST("/employee", EmployeeController.POST)
  //PUT
  r.PUT("/employee", EmployeeController.PUT)
  //DELETE
  r.DELETE("/employee", EmployeeController.DELETE)

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}