package xtypes

import (
	"encoding/json"
)

func Copy(dst, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dst)
}
