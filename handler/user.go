package handler

import (
	"net/http"

	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
		
	}
}
