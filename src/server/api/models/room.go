package models

import "github.com/jinzhu/gorm"

type Room struct {
  gorm.Model
  Owner_id uint `json:"owner_id"`
  Name string `json:"name"`
}

type Rooms []Room

func (room *Room) columns() ([]string) {
  columns := []string{"id", "name", "owner_id"}
  return columns
}

func (rooms *Rooms) All()  {
  DB.Select(new(Room).columns()).Find(&rooms)
}

func (room *Room) Find(id string)  {
  DB.Select(new(Room).columns()).First(&room, id)
}

func (room *Room) Messages(messages *Messages) {
  DB.Model(room).Related(messages)
}
