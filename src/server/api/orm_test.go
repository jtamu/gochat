package main

import (
  "fmt"
  _ "github.com/gin-gonic/gin"
  . "server/api/models"
)

func main()  {
  users := []User{}
  DB.Find(&users)
  for _, user := range users {
    fmt.Println(user.Name)
  }
}
