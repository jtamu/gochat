package main

import (
  "fmt"
  _ "github.com/lib/pq"
  "server/api/models"
)

func main()  {
  fmt.Println("auto-migration start")
  models.AutoMigrate()
  fmt.Println("auto-migration end")
}
