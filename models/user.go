package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UUID      string     `json:"uuid"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	DietID    *uint
	Diet      *Diet
}

func (u *User) TableName() string {
	return "users"
}

func (m Model) AddUser(user User) (*User, error) {
	var check User
	err := m.DB.Where(&User{UUID: user.UUID}).First(&check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if check.UUID != "" {
		return nil, errors.New("The user is already registered")
	}
	err = m.DB.Create(&user).Error
	return &user, err
}
func (m Model) UpdateUser(userID uint, dietID uint) error {
	return m.DB.Model(User{ID: userID}).Update("diet_id", dietID).Error
}
