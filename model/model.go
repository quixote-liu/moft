package model

import "github.com/quixote-liu/config"

var CONF = config.CONF()

type H map[string]interface{}

const (
	DirFile  = "./static_file"
	DirPhoto = "./static_photo"
)
