package controllers

import (
  "server/api/models"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "log"
  "encoding/json"
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
    message := new(models.Message)
    json.Unmarshal(msg, message)
    go message.Create()
    res, _ := json.Marshal(message)
    ctr.melody.BroadcastFilter(res, func(q *melody.Session) bool {
      // 同じURL => 同じRoomにいる人全員に送信する
      return q.Request.URL.Path == s.Request.URL.Path
    })
  })
}

// GET "/" : DBに保存されたMessageを全て取得

// GET "/ws" : websocket connection open
func (ctr *MessagesController) WsConnect(c *gin.Context) {
  ctr.melody.HandleRequest(c.Writer, c.Request)
}
