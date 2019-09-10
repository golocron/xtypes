package xtypes

import (
	"reflect"
	"testing"
)

type T1 struct {
	Name string
}

func TestCopy(t *testing.T) {
	type testCase struct {
		name string

		src *T1
		dst *T1
	}

	tests := []*testCase{
		&testCase{
			name: "VALID - Regular Struct",

			src: &T1{Name: "rainbow"},
			dst: &T1{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := Copy(tc.dst, tc.src); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.src, tc.dst) {
				t.Fatal("Not Equal: reflect")
			}
		})
	}
}

// Benchmarks
func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(&T1{Name: "rainbow"}, &T1{})
	}
}

//
