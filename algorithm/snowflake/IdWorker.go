package snowflake

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	twepoch            = int64(1288834974657)
	workerIdBits       = uint(5)
	datacenterIdBits   = uint(5)
	maxWorkerId        = -1 ^ (-1 << workerIdBits)
	maxDatacenterId    = -1 ^ (-1 << datacenterIdBits)
	sequenceBits       = uint(12)
	workerIdShift      = sequenceBits
	datacenterIdShift  = sequenceBits + workerIdBits
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits
	sequenceMask       = -1 ^ (-1 << sequenceBits)
	maxNextIdsNum      = 100
)

type IdWorker struct {
	sequence      int64
	lastTimestamp int64
	workerId      int64
	twepoch       int64
	datacenterId  int64
	mutex         sync.Mutex
}

func NewIdWorker(workerId, datacenterId int64, twepoch int64) (*IdWorker, error) {

	idWorker := &IdWorker{}
	if workerId > maxWorkerId || workerId < 0 {
		return nil, errors.New(fmt.Sprintf("worker Id: %d error", workerId))
	}

	if datacenterId > maxDatacenterId || datacenterId < 0 {
		return nil, errors.New(fmt.Sprintf("datacenter Id: %d error", datacenterId))
	}

	idWorker.workerId = workerId
	idWorker.datacenterId = datacenterId
	idWorker.lastTimestamp = -1
	idWorker.sequence = 0
	idWorker.twepoch = twepoch
	idWorker.mutex = sync.Mutex{}

	return idWorker, nil

}

func timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func tilNextMillis(lastTimestamp int64) int64 {
	timestamp := timeGen()
	for timestamp <= lastTimestamp {
		timestamp = timeGen()
	}
	return timestamp

}

func (id *IdWorker) NextId() (int64, error) {
	id.mutex.Lock()

	defer id.mutex.Unlock()

	timestamp := timeGen()
	if timestamp < id.lastTimestamp {
		return 0, errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", id.lastTimestamp-timestamp))
	}

	if id.lastTimestamp == timestamp {
		id.sequence = (id.sequence + 1) & sequenceMask
		if id.sequence == 0 {
			timestamp = tilNextMillis(id.lastTimestamp)
		}
	} else {
		id.sequence = 0
	}

	id.lastTimestamp = timestamp
	return ((timestamp - id.twepoch) << timestampLeftShift) | (id.datacenterId << datacenterIdShift) | (id.workerId << workerIdShift) | id.sequence, nil
}

func (id *IdWorker) NextIds(num int) ([]int64, error) {

	if num > maxNextIdsNum || num < 0 {
		return nil, errors.New(fmt.Sprintf("NextIds num: %d error", num))
	}

	ids := make([]int64, num)
	id.mutex.Lock()

	defer id.mutex.Unlock()

	for i := 0; i < num; i++ {
		timestamp := timeGen()
		if timestamp < id.lastTimestamp {
			return nil, errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", id.lastTimestamp-timestamp))
		}

		if id.lastTimestamp == timestamp {
			id.sequence = (id.sequence + 1) & sequenceMask
			if id.sequence == 0 {
				timestamp = tilNextMillis(id.lastTimestamp)
			}
		} else {
			id.sequence = 0
		}
		id.lastTimestamp = timestamp
		ids[i] = ((timestamp - id.twepoch) << timestampLeftShift) | (id.datacenterId << datacenterIdShift) | (id.workerId << workerIdShift) | id.sequence
	}
	return ids, nil
}
