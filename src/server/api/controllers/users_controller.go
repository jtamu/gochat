package controllers

import (
  . "server/api/models"
)

type UsersController struct{}

func (uc *UsersController) Index() (users []User) {
  All(&users)
  return
}

func (uc *UsersController) Show(id uint) (user User) {
  Find(&user, id)
  return
}
