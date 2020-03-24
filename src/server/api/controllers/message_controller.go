package controllers

import (
  "server/api/models"
)

type MessagesController struct{
  Controller
}

func NewMessagesController() (mc MessagesController) {
  mc.resource = new(models.Message)
  mc.resources = new(models.Messages)
  return
}
