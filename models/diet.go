package models

type Diet struct {
	ID          uint     `json:"id"`
	Type        dietType `json:"dietType"`
	Description string   `json:"description"`
	CelebrityID uint     `json:"celeb_id"`
	Meals       []Meal   `json:"meals"`
}

func (u *Diet) TableName() string {
	return "diet"
}

type Meal struct {
	ID     uint   `json:"-"`
	Name   string `json:"name"`
	DietID uint   `json:"-"`
	Foods  []Food `json:"foods"`
}

func (u *Meal) TableName() string {
	return "meal"
}

type Food struct {
	ID       uint    `json:"-"`
	Name     string  `json:"name"`
	Calories float64 `json:"calories"`
	MealID   uint    `json:"-"`
}

func (u *Food) TableName() string {
	return "food"
}

func (m *Model) GetDiet(id uint) (Diet, error) {
	diet := Diet{ID: id}
	err := m.DB.Debug().Preload("Meals.Foods").First(&diet).Error
	return diet, err
}
