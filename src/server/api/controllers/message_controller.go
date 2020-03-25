package controllers

import (
  "server/api/models"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "fmt"
  "encoding/json"
  "io/ioutil"
)

type MessagesController struct{
  Controller
  melody *melody.Melody
}

func NewMessagesController() (mc MessagesController) {
  mc.resource = new(models.Message)
  mc.resources = new(models.Messages)
  mc.melody = melody.New()
  mc.handleMessage()
  return
}

func (ctr *MessagesController) Melody() *melody.Melody {
  return ctr.melody
}

// Index : DBに保存されたMessageを全て取得

// HandleMessage : Messageを受信する
func (ctr *MessagesController) handleMessage() {
  ctr.melody.HandleMessage(func(s *melody.Session, msg []byte) {
    ctr.melody.Broadcast(msg)
  })
}

// Create : Messageを送信する
// @override
func (ctr *MessagesController) Create(c *gin.Context) {
  body, _ := ioutil.ReadAll(c.Request.Body)
  json_bytes := ([]byte)(body)
  msg := new(models.Message)
  json.Unmarshal(json_bytes, msg)
  fmt.Fprintln(c.Writer, msg)
}
