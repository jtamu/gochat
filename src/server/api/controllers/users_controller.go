package controllers

import (
  "server/api/models"
)

type UsersController struct{
  Controller
}

func NewUsersController() (uc UsersController) {
  uc.resource = new(models.User)
  uc.resources = new(models.Users)
  return
}
