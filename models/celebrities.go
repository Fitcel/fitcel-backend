package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Celebrity struct {
	ID       uint     `json:"id" gorm:"primarykey; not null"`
	Name     string   `json:"name" gorm:"column:name"`
	Avatar   string   `json:"avatar" gorm:"column:avatar; unique"`
	DietType dietType `json:"diet_type" gorm:"index"`
	DType    Diet     `json:"-" gorm:"foreignKey:Type;references:diet_type"`
	Diet     Diet     `json:"celebrityDiet"`
}

func (u *Celebrity) TableName() string {
	return "celebrities"
}

func (m *Model) AddCelebrity(celeb *Celebrity, avatarURL string) error {
	var check Celebrity
	err := m.DB.Where(&Celebrity{Name: celeb.Name}).First(&check).Error
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

func (m *Model) GetCelebrities() ([]Celebrity, error) {
	var celebs []Celebrity
	// err := m.DB.Find(&celebs).Error
	err := m.DB.Model(new(Celebrity)).Find(&celebs).Error

	return celebs, err
}

func (m *Model) GetCelebrity(id uint) (Celebrity, error) {
	celeb := Celebrity{ID: id}
	err := m.DB.Debug().First(&celeb).Error
	return celeb, err
}

func (m *Model) GetCelebritybyDietID(dietID uint) (Celebrity, error) {
	celeb := Celebrity{}
	diet, err := m.GetDiet(dietID)
	if err != nil {
		return Celebrity{}, err
	}
	err = m.DB.Debug().Where(&Celebrity{ID: diet.CelebrityID}).First(&celeb).Error
	return celeb, err
}
