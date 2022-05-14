package handler

import (
	"bufio"
	"io"
	"log"
	"moft/model"
	"moft/util"
	"net/http"
	"os"
)

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		util.JSONResponse(w, http.StatusBadRequest, model.H{
			"error":   err.Error(),
			"message": "parse request form failed",
		})
		return
	}

	// read file.
	f, fh, err := r.FormFile("file")
	if err != nil {
		util.JSONResponse(w, http.StatusBadRequest, model.H{
			"error":   err.Error(),
			"message": "get form file failed",
		})
		return
	}
	defer f.Close()
	// read file content to buffer.
	reader := bufio.NewReader(f)

	fileName := absoluteFileName(fh.Filename)

	// create file.
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("create file(%s) failed: %v", fileName, err)
		util.JSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	// wirte data.
	_, err = io.Copy(file, reader)
	if err != nil {
		log.Printf("write data to file(%s) failed: %v", fileName, err)
		util.JSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(200)
}
