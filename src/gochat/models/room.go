package models

type Room struct {
  Id uint `json:"id"`
  Owner_id uint `json:"owner_id"`
  Name string `json:"name"`
}
