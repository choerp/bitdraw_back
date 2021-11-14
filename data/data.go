package data

import (
	"encoding/json"
)

type UserData struct {
	Index     uint64
	LowPrice  float64
	HighPrice float64
}

func CreatedUserData(doc string) *[]UserData {
	var data []UserData
	json.Unmarshal([]byte(doc), &data)
	return &data
}
