package main

import (
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"

  "github.com/projects/app/controller"
)

func main() {
  r := gin.Default()

  r.Use(cors.Default())

  r.GET("/", func(c *gin.Context) {
    c.String(200, "OK")
  })
  r.GET("/users", controller.GetUsers)
  r.POST("/users", controller.CreateUser)
  r.GET("/users/:id", controller.GetUser)
  r.PUT("/users/:id", controller.UpdateUser)
  r.DELETE("/users/:id", controller.DeleteUser)
  r.GET("/users/:id/ws", controller.GetUserWebSocket)

  r.Run(":8080")
}
