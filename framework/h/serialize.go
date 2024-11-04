package h

import (
	"encoding/json"
)

// JsonSerializeOrEmpty serializes the given data as JSON, or returns an empty string if the serialization fails.
func JsonSerializeOrEmpty(data any) string {
	if data == nil {
		return ""
	}
	serialized, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(serialized)
}
