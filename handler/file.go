package handler

import (
	"bufio"
	"io"
	"log"
	"moft/model"
	"moft/util"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"error":   err.Error(),
			"message": "parse request form failed",
		})
		return
	}

	// read file.
	f, fh, err := r.FormFile("file")
	if err != nil {
		util.ResponseJSONErr(w, http.StatusBadRequest, model.H{
			"error":   err.Error(),
			"message": "get form file failed",
		})
		return
	}
	defer f.Close()
	// read file content to buffer.
	reader := bufio.NewReader(f)

	fileName := makeFileName(fh.Filename)

	// create file.
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("create file(%s) failed: %v", fileName, err)
		util.ResponseJSONErr(w, http.StatusInternalServerError, nil)
		return
	}

	// wirte data.
	_, err = io.Copy(file, reader)
	if err != nil {
		log.Printf("write data to file(%s) failed: %v", fileName, err)
		util.ResponseJSONErr(w, http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(200)
}

func makeFileName(fileName string) string {
	ext := filepath.Ext(fileName)
	fileName = strings.TrimSuffix(fileName, ext)

	name := fileName + "_" + time.Now().Format(time.RFC3339) + ext
	dir := model.DirFile
	dir = filepath.Clean(dir)
	if !strings.HasPrefix(dir, "/") {
		dir += "/"
	}
	return dir + name
}
