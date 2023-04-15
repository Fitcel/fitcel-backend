package controllers

import (
	"fitcel-backend/models"
	"strconv"
)

func (c *Controller) AddUser(user models.User) (*models.User, error) {
	return c.Model.AddUser(user)
}

func (c *Controller) UpdateUser(UUID string, dietID string) error {
	dietid, err := strconv.Atoi(dietID)
	if err != nil {
		return err
	}
	return c.Model.UpdateUser(UUID, uint(dietid))
}
