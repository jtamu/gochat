package controllers

import (
  "github.com/gin-gonic/gin"
)

var (
  uc = NewUsersController()
  rc = NewRoomsController()
  mc = NewMessagesController()
)

func Init() {
  r := gin.Default()

  ug := r.Group("/users")
  {
    ug.GET("", uc.Index)
    ug.GET("/:id", uc.Show)
  }

  rg := r.Group("/rooms")
  {
    rg.GET("", rc.Index)
    rg.GET("/:id", rc.Show)
  }

  mg := r.Group("/rooms/:id/messages")
  {
    mg.GET("", mc.Index)
    mg.GET("/ws", mc.WsConnect)
  }

  r.Run(":8080")
}
