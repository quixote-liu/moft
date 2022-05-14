package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Appendix struct {
	ID     string   `gorm:"primaryKey"`
	UserID string   `gorm:"not null"`
	Type   FileType `gorm:"not null"`
	Path   string   `gorm:"not null"`
	Name   string   `gorm:"not null"`
}

func CreateAppendix(db *gorm.DB, a Appendix) error {
	if a.ID == "" {
		a.ID = uuid.NewString()
	}
	return db.Create(&a).Error
}

func DeleteAppendix(db *gorm.DB, id string) error {
	return db.Where("id = ?").Delete(&Appendix{}).Error
}

func FindAppendix(db *gorm.DB, id string) (Appendix, error) {
	a := Appendix{}
	err := db.Where("id = ?").Find(&a).Error
	return a, err
}
