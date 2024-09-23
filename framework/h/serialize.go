package h

import (
	"encoding/json"
)

func JsonSerializeOrEmpty(data any) string {
	serialized, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(serialized)
}
