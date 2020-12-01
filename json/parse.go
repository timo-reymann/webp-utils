package json

import (
	"encoding/json"
)

func Parse(data []byte) (map[string]interface{}, error) {
	var parsed map[string]interface{}
	err := json.Unmarshal(data, &parsed)
	return parsed, err
}
