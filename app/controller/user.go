package controller

import (
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "github.com/projects/app/domain"
)

func GetUsers(c *gin.Context) {
  c.JSON(http.StatusOK, domain.GetUsers())
}

func GetUser(c *gin.Context) {
  id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
  c.JSON(http.StatusOK, domain.GetUser(id))
}

func CreateUser(c *gin.Context) {
  data := domain.UserDomain{}

  if err := c.BindJSON(&data); err != nil {
    c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
    return
  }

  data.Save()
  c.JSON(http.StatusOK, data)
}

func UpdateUser(c *gin.Context) {
  id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
  data := domain.GetUser(id)

  if err := c.BindJSON(&data); err != nil {
    c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
    return
  }

  data.Save()
  c.JSON(http.StatusOK, data)
}

func DeleteUser(c *gin.Context) {
  id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
  data := domain.GetUser(id)
  data.Delete()

  c.JSON(http.StatusOK, map[string]uint{"id": data.ID})
}

var MConnections = map[uint]*melody.Melody{}
func GetUserWebSocket(c *gin.Context) {
  id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
  data := domain.GetUser(id)

  mc, ok := MConnections[data.ID]
  if ok {

  } else {
    mc = melody.New()
    MConnections[data.ID] = mc

    mc.HandleMessage(func(s *melody.Session, msg []byte) {
      mc.Broadcast(msg)
    })
  }

  mc.HandleRequest(c.Writer, c.Request)
}
