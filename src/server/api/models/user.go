package models

import "github.com/jinzhu/gorm"

type User struct {
  gorm.Model
  Name string `json:"name"`
}

type Users []User

func (user *User) columns() ([]string) {
  columns := []string{"id", "name"}
  return columns
}

func (users *Users) All()  {
  DB.Select(new(User).columns()).Find(&users)
}

func (user *User) Find(id string)  {
  DB.Select(new(User).columns()).First(&user, id)
}
