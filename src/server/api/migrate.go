package main

import (
  "fmt"
  _ "github.com/lib/pq"
  . "server/api/models"
)

func main()  {
  fmt.Println("auto-migration start")
  for _, model := range Models {
    DB.AutoMigrate(model)
  }
  fmt.Println("auto-migration end")
}
