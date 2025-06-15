package server

import (
	"bike_store/database/models"
	"time"
)

var defaultSessionTime = 5 * time.Minute

type session struct {
	user    *models.User
	expires time.Time
}

func newExpiryTime() time.Time {
	return time.Now().Add(defaultSessionTime)
}

func newSession(user *models.User) *session {
	return &session{
		user:    user,
		expires: newExpiryTime(),
	}
}

func (s *session) Refresh() {
	s.expires = newExpiryTime()
}

func (s *session) HasExpired() bool {
	return time.Now().After(s.expires)
}
