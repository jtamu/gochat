package controllers

import (
  "server/api/models"
  "github.com/gin-gonic/gin"
)

type RoomsController struct{}

func (rc *RoomsController) Index(c *gin.Context) {
  rooms := models.Rooms{}
  rooms.All()
  c.JSON(200, rooms)
  return
}

func (rc *RoomsController) Show(c *gin.Context) {
  room := models.Room{}
  id := c.Params.ByName("id")
  room.Find(id)
  c.JSON(200, room)
  return
}
