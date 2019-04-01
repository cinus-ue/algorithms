package tokenbucket

import (
	"sync"
	"time"
)

type Bucket struct {
	capacity  int64
	tokens    chan struct{}
	rate      time.Duration
	rateMutex sync.Mutex
}

func NewBucket(rate time.Duration, capacity int64) *Bucket {

	tokens := make(chan struct{}, capacity)
	b := &Bucket{capacity, tokens, rate, sync.Mutex{}}

	go func(b *Bucket) {
		ticker := time.NewTicker(rate)
		for range ticker.C {
			b.tokens <- struct{}{}
		}
	}(b)
	return b
}

func (b *Bucket) GetRate() time.Duration {
	b.rateMutex.Lock()
	tmp := b.rate
	b.rateMutex.Unlock()
	return tmp
}

func (b *Bucket) SetRate(rate time.Duration) {
	b.rateMutex.Lock()
	b.rate = rate
	b.rateMutex.Unlock()
}

func (b *Bucket) AddToken(n int64) {

}

func (b *Bucket) withdrawTokens(n int64) error {
	for i := int64(0); i < n; i++ {
		<-b.tokens
	}
	return nil
}

func (b *Bucket) SpendToken(n int64) <-chan error {

	if n < 0 {
		n = 1
	}
	c := make(chan error)
	go func(b *Bucket, n int64, c chan error) {
		c <- b.withdrawTokens(n)
		close(c)
		return
	}(b, n, c)
	return c
}

func (b *Bucket) Drain() error {
	// TODO replace this with a more solid approach (such as replacing the channel altogether)
	for {
		select {
		case _ = <-b.tokens:
			continue
		default:
			return nil
		}
	}
}
