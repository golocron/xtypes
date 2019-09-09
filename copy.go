package xtypes

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func Copy(dst, src interface{}) error {
	buf := &bytes.Buffer{}

	enc := gob.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return err
	}

	dec := gob.NewDecoder(buf)

	return dec.Decode(dst)
}

func Copy2(dst, src interface{}) error {
	buf := &bytes.Buffer{}

	enc := json.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return err
	}

	dec := json.NewDecoder(buf)

	return dec.Decode(dst)
}
