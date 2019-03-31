package leakybucket

import (
	"sync"
	"time"
)

type bucket struct {
	capacity  uint
	remaining uint
	reset     time.Time
	rate      time.Duration
	mutex     sync.Mutex
}

func (b *bucket) Capacity() uint {
	return b.capacity
}


func (b *bucket) Remaining() uint {
	return b.remaining
}


func (b *bucket) Reset() time.Time {
	b.remaining = b.capacity
	return b.reset
}


func (b *bucket) Add(amount uint) (BucketState, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if time.Now().After(b.reset) {
		b.reset = time.Now().Add(b.rate)
		b.remaining = b.capacity
	}

	if amount > b.remaining {
		return BucketState{Capacity: b.capacity, Remaining: b.remaining, Reset: b.reset}, ErrorFull
	}

	b.remaining -= amount
	return BucketState{Capacity: b.capacity, Remaining: b.remaining, Reset: b.reset}, nil

}


type Storage struct {
	buckets map[string]*bucket
}


func New() *Storage {
	return &Storage{
		buckets: make(map[string]*bucket),
	}
}


func (s *Storage) Create(name string, capacity uint, rate time.Duration) (BucketI, error) {
	b, ok := s.buckets[name]
	if ok {
		return b, nil
	}

	b = &bucket{
		capacity:  capacity,
		remaining: capacity,
		reset:     time.Now().Add(rate),
		rate:      rate,
	}

	s.buckets[name] = b
	return b, nil
}
