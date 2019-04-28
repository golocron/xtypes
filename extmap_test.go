package xtypes

import (
	"bytes"
	"math"
	"testing"
)

func TestExtMap(t *testing.T) {
	m := make(ExtMap)
	if m == nil {
		t.Fatal("failed to create map")
	}
}

func TestExtMap_Get(t *testing.T) {
	m := make(ExtMap)

	key, value := "testGet", "testValue"
	m.Set(key, value)

	if v := m.Get(key); v != value {
		t.Fatalf("expected %s, got %s type", value, v)
	}
}

func TestExtMap_GetNil(t *testing.T) {
	var m ExtMap

	key := "testCase"
	expected := interface{}(nil)

	if v := m.Get(key); v != nil {
		t.Fatalf("expected %s, got %s type", expected, v)
	}
}

func TestExtMap_Set(t *testing.T) {
	m := make(ExtMap)

	key, value := "testGet", "testValue"
	m.Set(key, value)

	if v := m.Get(key); v != value {
		t.Fatalf("expected %s, got %s type", value, v)
	}
}

func TestExtMap_Del(t *testing.T) {
	m := make(ExtMap)

	key, value := "testGet", "testValue"
	m.Set(key, value)

	m.Del(key)

	expected := interface{}(nil)

	if v := m.Get(key); v != expected {
		t.Fatalf("expected %s, got %s type", expected, v)
	}
}

func TestExtMap_Encode(t *testing.T) {
	m := make(ExtMap)

	key, value := "testGet", "testValue"
	m.Set(key, value)

	expected := []byte(`{"testGet":"testValue"}`)

	if v := m.Encode(); !bytes.Equal(expected, v) {
		t.Fatalf("expected %s, got %s", expected, v)
	}
}

func TestExtMap_EncodeNil(t *testing.T) {
	var m ExtMap

	expected := []byte(nil)

	if v := m.Encode(); !bytes.Equal(expected, v) {
		t.Fatalf("expected %s, got %s", expected, v)
	}
}

func TestExtMap_EncodeError(t *testing.T) {
	m := make(ExtMap)

	key, value := "testGet", math.NaN()
	m.Set(key, value)

	expected := []byte(nil)

	if v := m.Encode(); !bytes.Equal(expected, v) {
		t.Fatalf("expected %s, got %s", expected, v)
	}
}
