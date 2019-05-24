package controller

import (
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "github.com/projects/app/lib"
  "github.com/projects/app/model"
)

var db = lib.DbConnection()

func GetUsers(c *gin.Context) {
  users := []model.User{}
  db.Find(&users)
  c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
  user := model.User{}
  id := c.Param("id")

  db.Where("ID = ?", id).First(&user)
  c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
  data := model.User{}
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
}

func UpdateUser(c *gin.Context) {
  user := model.User{}
  id := c.Param("id")

  data := model.User{}
  if err := c.BindJSON(&data); err != nil {
    c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
    return
  }

  db.Where("ID = ?", id).First(&user).Updates(&data)
  c.JSON(http.StatusOK, data)
}

func DeleteUser(c *gin.Context) {
  user := model.User{}
  id := c.Param("id")

  db.Where("ID = ?", id).Delete(&user)
}

var MConnections = map[string]*melody.Melody{}
func GetUserWebSocket(c *gin.Context) {
  user := model.User{}
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
}
