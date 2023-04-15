package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FoodAPIResponse struct {
	Name                string  `json:"name"`
	Calories            float64 `json:"calories"`
	ServingSizeG        float64 `json:"serving_size_g"`
	FatTotalG           float64 `json:"fat_total_g"`
	FatSaturatedG       float64 `json:"fat_saturated_g"`
	ProteinG            float64 `json:"protein_g"`
	SodiumMg            int     `json:"sodium_mg"`
	PotassiumMg         int     `json:"potassium_mg"`
	CholesterolMg       int     `json:"cholesterol_mg"`
	CarbohydratesTotalG float64 `json:"carbohydrates_total_g"`
	FiberG              float64 `json:"fiber_g"`
	SugarG              float64 `json:"sugar_g"`
}

func (s Service) GetFoodNutrition(query string) (FoodAPIResponse, error) {
	url := s.Food_Api_URL + url.QueryEscape(query)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return FoodAPIResponse{}, err
	}
	request.Header.Add("X-Api-Key", s.Food_Api_KEY)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return FoodAPIResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return FoodAPIResponse{}, err
	}
	var response []FoodAPIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return FoodAPIResponse{}, err
	}
	return response[0], nil
}
