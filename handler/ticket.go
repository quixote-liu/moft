package handler

import (
	"moft/model"
	"moft/util"
	"net/http"

	"gorm.io/gorm"
)

type TicketHandler struct {
	db *gorm.DB
}

func NewTicketHandler(db *gorm.DB) *TicketHandler {
	return &TicketHandler{db: db}
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	// authenticate user.
	sess, err := model.NewSession(r)
	if err != nil {
		util.Status(w, http.StatusUnauthorized)
		return
	}

	// parse form.
	if err := r.ParseForm(); err != nil {
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"message": "parse form data failed",
			"error":   err.Error(),
		})
		return
	}

	ticket := model.Ticket{}

	// get ticket message.
	ticket.Message = r.PostFormValue("message")
}
