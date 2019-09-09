package xtypes

import (
	"reflect"
	"testing"
)

type T1 struct {
	Name string
}

func TestCopy(t *testing.T) {
	v1 := &T1{Name: "rainbow"}
	v2 := &T1{}

	if err := Copy(v2, v1); err != nil {
		t.Fatal(err)
	}

	if v1.Name != v2.Name {
		t.Fatalf("Not Equal: expected %s => actual %s\n", v1.Name, v2.Name)
	}

	if !reflect.DeepEqual(v1, v2) {
		t.Fatal("Not Equal: reflect")
	}
}

func TestCopy2(t *testing.T) {
	v1 := &T1{Name: "rainbow"}
	v2 := &T1{}

	if err := Copy2(v2, v1); err != nil {
		t.Fatal(err)
	}

	if v1.Name != v2.Name {
		t.Fatalf("Not Equal: expected %s => actual %s\n", v1.Name, v2.Name)
	}

	if !reflect.DeepEqual(v1, v2) {
		t.Fatal("Not Equal: reflect")
	}
}

// Benchmarks
func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(&T1{Name: "rainbow"}, &T1{})
	}
}

func BenchmarkCopy2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy2(&T1{Name: "rainbow"}, &T1{})
	}
}

//
