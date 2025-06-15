package dto

import "encoding/json"

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *LoginDto) AsJson() ([]byte, error) {
	return json.Marshal(s)
}

func (s *LoginDto) FromJson(data []byte) error {
	return json.Unmarshal(data, s)
}
