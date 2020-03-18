package main

import (
  "fmt"
  _ "github.com/lib/pq"
  "gochat/models"
)

func main()  {
  fmt.Println("auto-migration start")
  for _, model := range models.Migrations() {
    models.DB.AutoMigrate(model)
  }
  fmt.Println("auto-migration end")
}
