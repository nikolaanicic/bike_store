package dto

import "encoding/json"

type RentBikeDto struct {
	CitizenID string `json:"citizen_id"`
	BikeID    int    `json:"bike_id"`
	City      string `json:"city"`
}

func (s *RentBikeDto) AsJson() ([]byte, error) {
	return json.Marshal(s)
}

func (s *RentBikeDto) FromJson(data []byte) error {
	return json.Unmarshal(data, s)
}
