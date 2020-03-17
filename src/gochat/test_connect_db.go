package main

import (
  "os"
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

func gormConnect() *gorm.DB {
  dbms := "postgres"
  host := os.Getenv("POSTGRES_HOST")
  port := os.Getenv("POSTGRES_PORT")
  user := os.Getenv("POSTGRES_USER")
  dbname := os.Getenv("POSTGRES_DB")
  password := os.Getenv("POSTGRES_PASSWORD")

  connect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
  db, err := gorm.Open(dbms, connect)

  if err != nil {
    panic(err.Error())
  }
  fmt.Println("db connected: ", &db)
  return db
}

func main()  {
  db := gormConnect()
  defer db.Close()
  db.LogMode(true)
}
