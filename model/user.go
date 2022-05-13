package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(db *gorm.DB, u *User) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return db.Create(u).Error
}

func DeleteUser(db *gorm.DB, id string) error {
	return db.Delete(&User{}, id).Error
}

func UpdateUser(db *gorm.DB, u *User) error {
	return db.Model(&User{}).Updates(u).Error
}

func FindUser(db *gorm.DB, filters map[string]interface{}) ([]User, error) {
	uu := []User{}
	err := db.Where(filters).Find(&uu).Error
	return uu, err
}
