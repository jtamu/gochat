package controllers

import (
  "server/api/models"
  "github.com/gin-gonic/gin"
)

type RoomsController struct{}

func (rc *RoomsController) Index(c *gin.Context) {
  rooms := models.Rooms{}
  models.All(&rooms)
  c.JSON(200, rooms)
  return
}

func (rc *RoomsController) Show(c *gin.Context) {
  room := models.Room{}
  id := c.Params.ByName("id")
  models.Find(&room, id)
  c.JSON(200, room)
  return
}
