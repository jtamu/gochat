package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

func gormConnect() *gorm.DB {
  dbms := "postgres"
  host := ""
  port := ""
  user := "postgres"
  dbname := ""
  pass := "password"
  connect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password)
  db, err := gorm.open(dbms, connect)

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
