package handler

import (
	"errors"
	"fmt"
	"log"
	"moft/model"
	"moft/util"
	"net/http"

	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := util.BindingJSON(r, &user); err != nil {
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"error": "your register user error",
		})
		return
	}

	// validate register message.
	u := model.User{
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
	}
	if err := h.validateUser(u); err != nil {
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"error": err.Error(),
		})
		return
	}

	// create user.
	if err := model.CreateUser(h.db, &u); err != nil {
		log.Printf("create user failed: %v", err)
		util.Status(w, http.StatusInternalServerError)
		return
	}
	uu, err := model.FindUserByName(h.db, u.UserName)
	if err != nil {
		log.Printf("find created user failed: %v", err)
		util.Status(w, http.StatusInternalServerError)
		return
	}

	util.ResponseJSONErr(w, http.StatusCreated, model.H{
		"user":    uu,
		"message": "create user success",
	})
}

func (h *UserHandler) validateUser(user model.User) error {
	if user.UserName == "" {
		return fmt.Errorf("missing user name, please retype")
	}
	if user.Email == "" {
		return fmt.Errorf("missing user email, please retype")
	}
	if user.Password == "" {
		return fmt.Errorf("missing user password, please retype")
	}

	_, err := model.FindUserByName(h.db, user.UserName)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("user name is exist, please retype")
	}
	return nil
}
