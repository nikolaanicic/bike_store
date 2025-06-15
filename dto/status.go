package dto

import "encoding/json"

type Status struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewStatus(code int, message string) *Status {
	return &Status{Message: message, Code: code}
}

func (s *Status) AsJson() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Status) FromJson(data []byte) error {
	return json.Unmarshal(data, s)
}
