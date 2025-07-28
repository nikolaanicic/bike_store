package dto

import "encoding/json"

type AddBikeDto struct {
	Type string `json:"type"`
	City string `json:"city"`
}

func (ab *AddBikeDto) AsJson() ([]byte, error) {
	return json.Marshal(ab)
}

func (ab *AddBikeDto) FromJson(data []byte) error {
	return json.Unmarshal(data, ab)
}
