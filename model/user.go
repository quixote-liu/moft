package model

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string
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
	return db.Delete(&User{}, "id = ?", id).Error
}

func UpdateUser(db *gorm.DB, u *User) error {
	return db.Model(&User{}).Updates(u).Error
}

func ListUser(db *gorm.DB, filters map[string]interface{}) ([]User, error) {
	uu := []User{}
	err := db.Where(filters).Find(&uu).Error
	return uu, err
}

func FindUserByName(db *gorm.DB, name string) (User, error) {
	u := User{}

	logger := logger.New(log.Default(), logger.Config{
		IgnoreRecordNotFoundError: true,
	})
	err := db.Session(&gorm.Session{Logger: logger}).Where("user_name = ?", name).First(&u).Error
	return u, err
}
