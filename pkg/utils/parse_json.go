package utils

import (
	"encoding/json"
	"io"
)

// ParseJSON returns object from JSON
func ParseJSON[T interface{}](reader io.ReadCloser) (*T, error) {
	var jsonData T

	if err := json.NewDecoder(reader).Decode(&jsonData); err != nil {
		return nil, err
	}

	return &jsonData, nil
}
