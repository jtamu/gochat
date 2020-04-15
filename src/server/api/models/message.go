package models

import (
  "github.com/jinzhu/gorm"
  "server/api/lib"
)

type Message struct {
  gorm.Model
  RoomId uint `json:"room_id"`
  UserId uint `json:"user_id"`
  Content string `json:"content"`
}

type Messages []Message

func (message *Message) columns() ([]string) {
  columns := []string{"id", "content", "room_id", "user_id"}
  return columns
}

func (messages *Messages) All() {
  DB.Select(new(Message).columns()).Find(&messages)
}

func (message *Message) Find(id string) {
  DB.Select(new(Message).columns()).First(&message, id)
}

func (message *Message) User(user *User) {
  user.Find(lib.ToS(message.UserId))
}

func (message *Message) Room(room *Room)  {
  room.Find(lib.ToS(message.RoomId))
}

func (message *Message) Create() {
  DB.Create(message)
}
