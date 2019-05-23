package main

import (
  // "fmt"
  // "log"
  // "os"
  "time"
  "net/http"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/projects/app/models"
)

func main() {
  r := gin.Default()
  m := melody.New()

  r.Use(cors.Default())

  db, err := gorm.Open("postgres", "host=db user=root password=root dbname=chat sslmode=disable")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  db.AutoMigrate(models.User{})
  db.LogMode(true)

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

  r.POST("/users", func(c *gin.Context) {
    data := models.User{}
    now := time.Now()
    data.CreatedAt = now
    data.UpdatedAt = now

    if err := c.BindJSON(&data); err != nil {
      c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
      return
    }
    db.NewRecord(data)
    db.Create(&data)
    if db.NewRecord(data) == false {
      c.JSON(http.StatusOK, data)
    }
  })

  r.GET("/users", func(c *gin.Context) {
    users := []models.User{}
    db.Find(&users)
    c.JSON(http.StatusOK, users)
  })

  r.GET("/users/:id", func(c *gin.Context) {
    user := models.User{}
    id := c.Param("id")

    db.Where("ID = ?", id).First(&user)
    c.JSON(http.StatusOK, user)
  })

  MConnections := map[string]*melody.Melody{}
  r.GET("/users/:id/ws", func(c *gin.Context) {
    user := models.User{}
    id := c.Param("id")

    db.Where("ID = ?", id).First(&user)

    mc, ok := MConnections[id]
    if ok {

    } else {
      mc = melody.New()
      MConnections[id] = mc

      mc.HandleMessage(func(s *melody.Session, msg []byte) {
        mc.Broadcast(msg)
      })
    }

    mc.HandleRequest(c.Writer, c.Request)
  })

  r.PUT("/users/:id", func(c *gin.Context) {
    user := models.User{}
    id := c.Param("id")

    data := models.User{}
    if err := c.BindJSON(&data); err != nil {
      c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
    }

    db.Where("ID = ?", id).First(&user).Updates(&data)
  })

  r.DELETE("/users/:id", func(c *gin.Context) {
    user := models.User{}
    id := c.Param("id")

    db.Where("ID = ?", id).Delete(&user)
  })

  r.Run(":8080")
}
