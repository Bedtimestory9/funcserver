// Package session
package session

import (
	"net/http"
	"slices"
	"strings"
	"sync"
)

type SessionManager struct {
	mu          sync.Mutex
	authedUsers []string
}

type TMPLData struct {
	Title    string
	UserID   string
	IsAuthed bool
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		authedUsers: []string{},
	}
}

func (s *SessionManager) AddUserToSession(u string) {
	s.authedUsers = append(s.authedUsers, u)
}

func (s *SessionManager) CheckIfInSession(u string) bool {
	return slices.Contains(s.authedUsers, u)
}

func getUserID(r *http.Request) string {
	q := r.URL.String()
	s := strings.Split(q, "/")
	id := s[len(s)-1]
	return id
}

func pipeSessionData(authed bool, u string) TMPLData {
	var tplData TMPLData

	if authed {
		tplData = TMPLData{
			Title:    "User Session",
			UserID:   "User " + u + " logged in",
			IsAuthed: true,
		}
	} else {
		tplData = TMPLData{
			Title:    "Guest Session",
			UserID:   "Please Log In",
			IsAuthed: false,
		}
	}

	return tplData
}

func SessionPipe(w http.ResponseWriter, r *http.Request, s *SessionManager) TMPLData {
	userID := getUserID(r)

	isAuthed := s.CheckIfInSession(userID)

	return pipeSessionData(isAuthed, userID)
}
