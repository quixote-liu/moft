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

	// TODO: validate user email
	// TODO: validate email captcha

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
	if len(user.UserName) < 6 {
		return fmt.Errorf("the user name length must more than 6, please retype")
	}
	if user.Email == "" {
		return fmt.Errorf("missing user email, please retype")
	}
	if user.Password == "" {
		return fmt.Errorf("missing user password, please retype")
	}
	if len(user.Password) < 6 {
		return fmt.Errorf("the password length must more than 6, please retype")
	}

	_, err := model.FindUserByName(h.db, user.UserName)
	if err == nil {
		return fmt.Errorf("the user name is exist, please retype")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	log.Printf("find user from database by name %s failed: %v", user.UserName, err)
	return fmt.Errorf("system internal error, register user failed")
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := util.BindingJSON(r, &user); err != nil {
		util.ResponseJSONErr(w, http.StatusNotFound, model.H{
			"error":   "the request body message error",
			"message": err.Error(),
		})
		return
	}

	// find user.
	uu, err := model.ListUser(h.db, map[string]interface{}{
		"user_name": user.UserName,
		"password":  user.Password,
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"error": "user name or password error",
		})
		return
	} else if err != nil {
		log.Printf("find user by user name and password failed: %v", err)
		util.Status(w, http.StatusInternalServerError)
		return
	} else if len(uu) > 1 {
		log.Printf("[system internal error]: there are multiple identical accounts")
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"error": "the account is abnormal. Please contact the administrator",
		})
		return
	}

	u := uu[0]

	// build session.
	sess := model.NewSession(r)
	sess.SetValues(map[string]interface{}{
		"user_name": u.UserName,
		"user_id":   u.ID,
	})
	if err := sess.Save(w, r); err != nil {
		util.ResponseJSONErr(w, http.StatusInternalServerError, model.H{
			"error": fmt.Sprintf("build session failed: %v", err),
		})
		return
	}

	return
}
