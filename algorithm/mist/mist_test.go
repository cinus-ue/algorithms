package mist

import (
	"testing"
)

func TestID(t *testing.T) {
	mist := NewMist()
	var prev int64
	for i := 0; i < 10; i++ {
		id := mist.Generate()
		if id < prev {
			t.Error("generated value is lower than the previous value")
		}
		prev = id
	}
}

func BenchmarkID(t *testing.B) {
	mist := NewMist()
	for i := 0; i < t.N; i++ {
		mist.Generate()
	}
}
