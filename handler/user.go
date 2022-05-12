package handler

import (
	"net/http"

	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Optimize: .....
}
