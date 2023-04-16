package controllers

import (
	"encoding/json"
	"fitcel-backend/models"
	"mime/multipart"
	"strconv"
)

func (c *Controller) AddCeleb(jsonInput string, avatar *multipart.FileHeader) error {
	celeb := new(models.Celebrity)
	json.Unmarshal([]byte(jsonInput), celeb)
	celeb.DietType = celeb.Diet.Type

	// Store the avatar for celeb in cloud bucket
	avatarURL, err := c.storeinCloud(celeb.Name, avatar)
	if err != nil {
		return err
	}

	// Update Calories for foods
	c.updateCalories(celeb.Diet.Meals)

	// Add The data to database
	err = c.Model.AddCelebrity(celeb, avatarURL)
	if err != nil {
		return c.Services.DeleteAvatarinCloud(celeb.Name) // delete the stored image if data wasn't stored
	}
	return nil
}

func (c *Controller) GetCelebs() ([]models.Celebrity, error) {
	celebs, err := c.Model.GetCelebrities()
	if err != nil {
		return []models.Celebrity{}, err
	}
	return celebs, nil
}

func (c *Controller) GetCeleb(id string) (models.Celebrity, error) {
	celeb_id, err := strconv.Atoi(id)
	celeb, err := c.Model.GetCelebrity(uint(celeb_id))
	if err != nil {
		return models.Celebrity{}, err
	}
	return celeb, nil
}

func (c *Controller) GetCelebByDietID(dietID string) (models.Celebrity, error) {
	celeb_id, err := strconv.Atoi(dietID)
	celeb, err := c.Model.GetCelebritybyDietID(uint(celeb_id))
	if err != nil {
		return models.Celebrity{}, err
	}
	return celeb, nil
}
func (c *Controller) storeinCloud(name string, avatar *multipart.FileHeader) (string, error) {
	src, err := avatar.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	avatarURL, err := c.Services.StoreAvatarinCloud(name, src)
	if err != nil {
		return "", err
	}
	return avatarURL, nil
}

func (c *Controller) updateCalories(meals []models.Meal) {
	for _, meal := range meals {
		for i := range meal.Foods {
			resp, err := c.Services.GetFoodNutrition(meal.Foods[i].Name)
			if err != nil {
				continue // ignore the error and continue populatiing calories
			}
			meal.Foods[i].Calories = resp.Calories
		}
	}
}
