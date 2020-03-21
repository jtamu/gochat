package controllers

import (
  "server/api/models"
  "github.com/gin-gonic/gin"
)

type Controller struct{
  resource models.Resource
  resources models.Resources
}

func (ctr *Controller) Index(c *gin.Context) {
  resources := ctr.resources
  resources.All()
  c.JSON(200, resources)
  return
}

func (ctr *Controller) Show(c *gin.Context) {
  resource := ctr.resource
  id := c.Params.ByName("id")
  resource.Find(id)
  c.JSON(200, resource)
  return
}
