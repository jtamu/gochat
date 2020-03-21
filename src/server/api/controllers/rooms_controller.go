package controllers

import (
  "server/api/models"
)

type RoomsController struct{
  Controller
}

func NewRoomsController() (rc RoomsController) {
  rc.resource = new(models.Room)
  rc.resources = new(models.Rooms)
  return
}
