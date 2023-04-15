package controllers

import (
	"fitcel-backend/models"
	"strconv"
)

func (c *Controller) AddUser(user models.User) (*models.User, error) {
	return c.Model.AddUser(user)
}

func (c *Controller) UpdateUser(userID string, dietID string) error {
	userid, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}
	dietid, err := strconv.Atoi(dietID)
	if err != nil {
		return err
	}
	return c.Model.UpdateUser(uint(userid), uint(dietid))
}
