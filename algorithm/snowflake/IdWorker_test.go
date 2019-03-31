package snowflake

import (
	"fmt"
	"testing"
)

func TestID(t *testing.T) {

	id, err := NewIdWorker(0, 0, twepoch)
	if err != nil {
		t.Errorf("NewIdWorker(0, 0) error(%v)", err)
	}

	sid, err := id.NextId()
	if err != nil {
		t.Errorf("id.NextId() error(%v)", err)
	}

	fmt.Printf("snowflake id: %d \n", sid)

	sids, err := id.NextIds(10)
	if err != nil {
		t.Errorf("id.NextId() error(%v)", err)
	}

	fmt.Printf("snowflake ids: %v", sids)
}

func BenchmarkID(t *testing.B) {

	id, err := NewIdWorker(0, 0, twepoch)
	if err != nil {
		t.Errorf("NewIdWorker(0, 0) error(%v)", err)
	}

	for i := 0; i < t.N; i++ {
		if _, err := id.NextId(); err != nil {
			t.FailNow()
		}
	}

}
