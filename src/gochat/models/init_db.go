package models

import (
  "os"
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

var DB = Init()

func Init() *gorm.DB {
  dbms := "postgres"
  host := os.Getenv("POSTGRES_HOST")
  port := os.Getenv("POSTGRES_PORT")
  user := os.Getenv("POSTGRES_USER")
  dbname := os.Getenv("POSTGRES_DB")
  password := os.Getenv("POSTGRES_PASSWORD")

  connect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
  db, err := gorm.Open(dbms, connect)
  db.LogMode(true)

  if err != nil {
    panic(err.Error())
  }
  return db
}
