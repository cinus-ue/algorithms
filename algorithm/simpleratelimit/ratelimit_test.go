package simpleratelimit

import (
	"testing"
	"time"
)

func TestLimit(t *testing.T) {

	rl := New(0, 0*time.Second)
	if rl.Limit() != false {
		t.Error()
	}

	r2 := New(1, time.Second)
	for i := 0; i < 10; i++ {
		if i == 0 && r2.Limit() != false {
			t.Error()
		} else if r2.Limit() != true {
			t.Error()
		}
	}

	r3 := New(1, time.Second)
	for i := 0; i < 2; i++ {
		if i == 0 && r3.Limit() != false {
			t.Error()
		} else {
			r3.UpdateRate(2)
			if r3.Limit() != true {
				t.Error()
			}
		}
	}

	r4 := New(1, time.Second)
	for i := 0; i < 10; i++ {
		if i == 0 && r4.Limit() != false {
			t.Error()
		} else {
			r4.Undo()
			if r4.Limit() != false {
				t.Error()
			}
		}
	}

	for i := 0; i < 10; i++ {
		r4.Limit()
	}
	r4.Undo()

}

func BenchmarkLimit(b *testing.B) {

	rl := New(1, time.Second)
	for i := 0; i < b.N; i++ {
		rl.Limit()
	}
}
