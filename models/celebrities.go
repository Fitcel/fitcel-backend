package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}

type Celebrities struct {
	ID       uint     `json:"id" gorm:"primarykey; not null"`
	Name     string   `json:"name" gorm:"column:name"`
	Avatar   string   `json:"avatar" gorm:"column:avatar; unique"`
	DietType dietType `json:"diet_type" gorm:"index"`
	DType    Diet     `json:"-" gorm:"foreignKey:Type;references:diet_type"`
	Diet     Diet     `json:"celebrityDiet"`
}

func (u *Celebrities) TableName() string {
	return "celebrities"
}

func (m *Model) AddCelebrity(celeb *Celebrities, avatarURL string) error {
	var check Celebrities
	err := m.DB.Where(&Celebrities{Name: celeb.Name}).First(&check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if check.Name != "" {
		return errors.New("The celebrity already exists")
	}
	celeb.Avatar = avatarURL
	fmt.Println(celeb.Diet)
	err = m.DB.Debug().Create(celeb).Error
	return err
}

func (m *Model) GetCelebrities() ([]Celebrities, error) {
	var celebs []Celebrities
	// err := m.DB.Find(&celebs).Error
	err := m.DB.Model(new(Celebrities)).Find(&celebs).Error

	return celebs, err
}

func (m *Model) GetCelebrity(id uint) (Celebrities, error) {
	celeb := Celebrities{ID: id}
	err := m.DB.Debug().First(&celeb).Error
	return celeb, err
}
