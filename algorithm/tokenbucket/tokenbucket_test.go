package tokenbucket

import (
	"testing"
	"time"
)

func Example_BucketUse() {
	bucket := NewBucket(5*time.Second, 3)
	<-bucket.SpendToken(1)
	RegulatedAction()
}

func RegulatedAction() {
	// Some expensive action goes on here
}

func Test_BucketBuffering(t *testing.T) {

	const RATE = 4 * time.Second
	const CAPACITY = 3
	const ERROR = 500 * time.Millisecond
	b := NewBucket(RATE, CAPACITY)

	time.Sleep(CAPACITY * RATE)

	before := time.Now()

	<-b.SpendToken(1)
	<-b.SpendToken(1)
	<-b.SpendToken(1)

	after := time.Now()
	if diff := after.Sub(before); diff > RATE {
		t.Errorf("Waited %d seconds, though this should have been nearly instantaneous", diff)
	}
}

func Test_BucketCreation(t *testing.T) {

	const RATE = 4 * time.Second
	const CAPACITY = 3
	const ERROR = 500 * time.Millisecond
	const EXPECTED_DURATION = RATE * CAPACITY

	b := NewBucket(RATE, CAPACITY)

	<-b.SpendToken(1)
	<-b.SpendToken(1)
	<-b.SpendToken(1)
	<-b.SpendToken(1)

	before := time.Now()

	<-b.SpendToken(1)
	<-b.SpendToken(1)
	<-b.SpendToken(1)

	after := time.Now()

	lower := EXPECTED_DURATION - ERROR
	upper := EXPECTED_DURATION + ERROR

	if diff := after.Sub(before); diff < lower || diff > upper {
		t.Errorf("Waited %s seconds, though really should have waited between %s and %s", diff.String(), lower.String(), upper.String())
	}
}
