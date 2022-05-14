package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID         string
	PhotoPaths []string
	FilePaths  []string
	Message    string `gorm:"type:TEXT"`
	Status     TicketStatus
}

func CreateTicket(db *gorm.DB, t Ticket) error {
	if t.ID == "" {
		t.ID = uuid.NewString()
	}
	t.Status = TicketStatusPending
	return db.Create(&t).Error
}

func FindTicket(db *gorm.DB, id string) (Ticket, error) {
	t := Ticket{}
	err := db.Where("id = ?", id).First(&t).Error
	return t, err
}

func FindTicketByUserID(db *gorm.DB, userID string) (Ticket, error) {
	t := Ticket{}
	err := db.Where("user_id = ?", userID).First(&t).Error
	return t, err
}

func DeleteTicket(db *gorm.DB, id string) error {
	return db.Delete(&Ticket{}, id).Error
}
