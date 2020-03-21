package controllers

import (
  "github.com/gin-gonic/gin"
)

var (
  uc = new(UsersController)
)

func Init() {
  r := gin.Default()
  u := r.Group("/users")
  {
    u.GET("", uc.Index)
    u.GET("/:id", uc.Show)
  }
  r.Run(":8080")
}
