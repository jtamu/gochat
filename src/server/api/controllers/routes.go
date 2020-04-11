package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

var (
  uc = NewUsersController()
  rc = NewRoomsController()
  mc = NewMessagesController()
)

func Init() {
  // CORS 対応
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:3001"}

  r := gin.Default()
  r.Use(cors.New(config))

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

  r.Run(":80")
}
