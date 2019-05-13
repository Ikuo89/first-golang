package main

import (
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
  m := melody.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "index",
		})
	})

  r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
