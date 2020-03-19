package main

import (
  "fmt"
  "server/api/controllers"
)

var (
  uc = new(controllers.UsersController)
)

func main() {

  users := uc.Index()
  for _, user := range users {
    fmt.Println(user.Name)
  }
  user := uc.Show(2)
  fmt.Println(user.Name)
}
