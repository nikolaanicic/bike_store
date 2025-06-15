package dto

import "encoding/json"

type RegisterDto struct {
	CitizenID string `json:"citizen_id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

func (s *RegisterDto) AsJson() ([]byte, error) {
	return json.Marshal(s)
}

func (s *RegisterDto) FromJson(data []byte) error {
	return json.Unmarshal(data, s)
}
