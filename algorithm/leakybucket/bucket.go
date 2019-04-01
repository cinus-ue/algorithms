package leakybucket

import (
	"errors"
	"time"
)

var ErrorFull = errors.New("add exceeds free capacity")

type BucketState struct {
	Capacity  uint
	Remaining uint
	Reset     time.Time
}

type BucketI interface {
	Capacity() uint
	Remaining() uint
	Reset() time.Time
	Add(uint) (BucketState, error)
}

type StorageI interface {
	Create(name string, capacity uint, rate time.Duration) (BucketI, error)
}
