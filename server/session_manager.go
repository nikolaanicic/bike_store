package server

import (
	"bike_store/database/models"
	"crypto/sha256"
	"encoding/hex"
)

type SessionManager struct {
	sessions map[string]*session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{}
}

func getToken(username string, name string) string {
	data := make([]byte, 0)
	data = append(data, []byte(username)...)
	data = append(data, []byte(name)...)
	encoded := sha256.Sum256(data)

	return hex.EncodeToString(encoded[:])

}

func (s *SessionManager) StartSession(user *models.User) string {
	token := getToken(user.CitizenID, user.Name)
	s.sessions[token] = newSession(user)

	return token
}

func (s *SessionManager) ValidateSession(token string) bool {
	if session, ok := s.sessions[token]; ok {
		return !session.HasExpired()
	}

	return false
}

func (s *SessionManager) RemoveSession(token string) {
	delete(s.sessions, token)
}

func (s *SessionManager) Refresh(token string) {
	if session, ok := s.sessions[token]; ok {
		session.Refresh()
	}
}
