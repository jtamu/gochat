package models

import "github.com/jinzhu/gorm"

type Room struct {
  gorm.Model
  Owner_id uint `json:"owner_id"`
  Name string `json:"name"`
}
