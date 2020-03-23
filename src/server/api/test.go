package main

import (
  "server/api/models"
  "fmt"
)

func main() {
  room := new(models.Room)
  room.Find("1")
  fmt.Println(room)
  messages := new(models.Messages)
  room.Messages(messages)
  fmt.Println(messages)
}
