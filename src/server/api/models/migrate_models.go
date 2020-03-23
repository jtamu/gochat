package models

var models = []interface{}{
  &User{},
  &Room{},
  &Message{},
}

func AutoMigrate()  {
  for _, model := range models {
    DB.AutoMigrate(model)
  }
}
