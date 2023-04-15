package controllers

import (
	"fitcel-backend/models"
	"strconv"
)

func (c *Controller) GetCelebDiet(id string) (models.Diet, error) {
	dietId, err := strconv.Atoi(id)
	if err != nil {
		return models.Diet{}, err
	}
	diet, err := c.Model.GetDiet(uint(dietId))
	if err != nil {
		return models.Diet{}, err
	}
	return diet, nil
}
