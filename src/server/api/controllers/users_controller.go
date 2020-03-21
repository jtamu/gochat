package controllers

import (
  "server/api/models"
  "github.com/gin-gonic/gin"
)

type UsersController struct{}

func (uc *UsersController) Index(c *gin.Context) {
  users := models.Users{}
  models.All(&users)
  c.JSON(200, users)
  return
}

func (uc *UsersController) Show(c *gin.Context) {
  user := models.User{}
  id := c.Params.ByName("id")
  models.Find(&user, id)
  c.JSON(200, user)
  return
}
