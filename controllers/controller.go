package controllers

import (
	"fitcel-backend/models"
	"fitcel-backend/services"
)

type Controller struct {
	Model    models.Model
	Services services.Service
}
