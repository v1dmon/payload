package wgen

import (
	"encoding/json"
)

func marshal(v any) ([]byte, error) {
	enc, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return enc, nil
}
