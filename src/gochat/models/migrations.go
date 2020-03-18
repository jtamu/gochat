package models

func Migrations() []interface{} {
  models := []interface{}{
    &User{},
    &Room{},
  }
  return models
}
