package xtypes

import (
	"encoding/json"
)

// ExtMap represents an extended map with come useful methods for convenience.
//
// It almost the same as url.Values/http.Header.
// The main purpose of this type is to put values and then marshal them to JSON.
// This is NOT thread safe.
type ExtMap map[string]interface{}

// Get gets the value associated with the given key.
// If there is no value associated with the key, Get returns the nil.
func (m ExtMap) Get(key string) interface{} {
	if m == nil {
		return nil
	}

	return m[key]
}

// Set sets the key to value. It replaces the existing value.
func (m ExtMap) Set(key string, value interface{}) {
	m[key] = value
}

// Del deletes the value associated with key.
func (m ExtMap) Del(key string) {
	delete(m, key)
}

// Encode encodes the values into raw json bytes.
func (m ExtMap) Encode() []byte {
	if m == nil {
		return nil
	}

	raw, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	return raw
}
