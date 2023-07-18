package common

import (
	"encoding/json"
)

func Unmarshal(bytes []byte, v any) error {
	err := json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}
