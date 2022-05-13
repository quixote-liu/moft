package model

import (
	"moft/util"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(util.RandomInt32()))

type Session struct {
	*sessions.Session
}

func NewSession(r *http.Request) *Session {
	sess, _ := store.Get(r, "Moft-Session")
	return &Session{
		Session: sess,
	}
}

func (s *Session) SetValues(values map[string]interface{}) {
	for k, v := range values {
		s.Session.Values[k] = v
	}
}

func (s *Session) GetString(key string) string {
	vv, ok := s.Values[key]
	if !ok {
		return ""
	}
	v, ok := vv.(string)
	if !ok {
		return ""
	}
	return v
}

func (s *Session) Save(w http.ResponseWriter, r *http.Request) error {
	return s.Session.Save(r, w)
}
