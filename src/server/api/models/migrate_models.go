package models

var models = []interface{}{
  &User{},
  &Room{},
}

func AutoMigrate()  {
  for _, model := range models {
    DB.AutoMigrate(model)
  }
}
