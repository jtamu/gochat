package controllers

import (
  "server/api/models"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "log"
  _ "encoding/json"
  _ "io/ioutil"
)

type MessagesController struct{
  Controller
  melody *melody.Melody
}

func NewMessagesController() (mc MessagesController) {
  mc.resource = new(models.Message)
  mc.resources = new(models.Messages)
  mc.melody = melody.New()
  mc.init_melody()
  return
}

func (ctr *MessagesController) init_melody() {

  ctr.melody.HandleConnect(func(s *melody.Session) {
    log.Printf("websocket connection open. [session: %#v]\n", s)
  })

  ctr.melody.HandleDisconnect(func(s *melody.Session) {
    log.Printf("websocket connection open. [session: %#v]\n", s)
  })

  ctr.melody.HandleMessage(func(s *melody.Session, msg []byte) {
    ctr.melody.Broadcast(msg)
  })
}

// GET "/" : DBに保存されたMessageを全て取得

// GET "/ws" : websocket connection open
func (ctr *MessagesController) WsConnect(c *gin.Context)  {
  ctr.melody.HandleRequest(c.Writer, c.Request)
}

/*
// Create : Messageを送信する
// @override
func (ctr *MessagesController) Create(c *gin.Context) {
  body, _ := ioutil.ReadAll(c.Request.Body)
  json_bytes := ([]byte)(body)
  msg := new(models.Message)
  json.Unmarshal(json_bytes, msg)
  c.JSON(200, msg)
}
*/
