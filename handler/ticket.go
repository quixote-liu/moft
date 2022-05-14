package handler

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"moft/model"
	"moft/util"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
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
	userID := sess.GetString("user_id")
	if userID == "" {
		log.Printf("[error]: the user id of session is empty")
		util.Status(w, http.StatusInternalServerError)
		return
	}

	// parse form.
	if err := r.ParseForm(); err != nil {
		util.JSONResponse(w, http.StatusBadRequest, model.H{
			"message": "parse form data failed",
			"error":   err.Error(),
		})
		return
	}

	type file struct {
		ftype string
		path  string
		buf   *bufio.Reader
	}
	files := make([]file, 0)
	defer func() {
		for _, f := range files {
			if f.buf != nil {
				f.buf.Reset(nil)
			}
		}
	}()

	// read files content.
	for k := range r.PostForm {
		var f multipart.File
		var fh *multipart.FileHeader
		var err error
		var ftype string

		switch {
		case strings.HasPrefix(k, "file_"):
			f, fh, err = r.FormFile(k)
			ftype = "file"
		case strings.HasPrefix(k, "photo_"):
			f, fh, err = r.FormFile(k)
			ftype = "photo"
		default:
			continue
		}
		if err != nil {
			util.JSONResponse(w, http.StatusBadRequest, model.H{
				"error": fmt.Sprintf("read file content failed: %v", err),
			})
			return
		}

		files = append(files, file{
			ftype: ftype,
			path:  absoluteFileName(fh.Filename),
			buf:   bufio.NewReader(f),
		})
	}

	ticket := model.Ticket{
		UserID:     userID,
		PhotoPaths: make([]string, 0),
		FilePaths:  make([]string, 0),
		Message:    r.PostFormValue("message"),
	}
	for _, f := range files {
		ff, err := os.Create(f.path)
		if err != nil {
			log.Printf("create file with path %s failed: %v", f.path, err)
			continue
		}
		_, err = io.Copy(ff, f.buf)
		if err != nil {
			log.Printf("write file content to target file failed: %v", err)
			continue
		}

		switch f.ftype {
		case "file":
			ticket.FilePaths = append(ticket.FilePaths, f.path)
		case "photo":
			ticket.PhotoPaths = append(ticket.PhotoPaths, f.path)
		}
	}

	// create ticket.
	if err := model.CreateTicket(h.db, ticket); err != nil {
		log.Printf("insert ticket secord into database failed: %v", err)
		util.Status(w, http.StatusInternalServerError)
		return
	}

	// response.
	util.JSONResponse(w, http.StatusCreated, model.H{
		"message":   "SUCCESS",
		"ticket_id": ticket.ID,
	})
}

func absoluteFileName(fileName string) string {
	name := uuid.NewString() + filepath.Ext(fileName)
	dir := filepath.Clean(model.DirFile)
	if !strings.HasPrefix(dir, "/") {
		dir += "/"
	}
	return dir + name
}
