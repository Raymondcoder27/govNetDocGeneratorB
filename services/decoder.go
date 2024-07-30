package services

import (
	"encoding/json"
)


func DecodeJSON(jsonData string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	return data, err
}
