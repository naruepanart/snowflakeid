package main

import (
	"sync"
	"time"
)

const (
	timestampBits  = 41
	machineIDBits  = 10
	sequenceBits   = 12
	maxMachineID   = -1 ^ (-1 << machineIDBits)
	maxSequenceNum = -1 ^ (-1 << sequenceBits)
)

// Snowflake is a structure representing a unique ID generator.
type Snowflake struct {
	mu          sync.Mutex
	lastTime    int64
	machineID   int64
	sequenceNum int64
}

// NewSnowflake creates a new Snowflake instance with the given machine ID.
func NewSnowflake(machineID int64) *Snowflake {
	if machineID < 0 || machineID > maxMachineID {
		panic("Invalid machine ID")
	}

	return &Snowflake{
		machineID: machineID,
	}
}

// GenerateID generates a unique ID.
func (s *Snowflake) GenerateID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentTime := time.Now().UnixNano() / 1e6

	if currentTime == s.lastTime {
		s.sequenceNum = (s.sequenceNum + 1) & maxSequenceNum
		if s.sequenceNum == 0 {
			currentTime = s.waitNextMillis()
		}
	} else {
		s.sequenceNum = 0
	}

	s.lastTime = currentTime

	id := (currentTime << (machineIDBits + sequenceBits)) |
		(s.machineID << sequenceBits) |
		s.sequenceNum

	return id
}

// waitNextMillis waits until the next millisecond.
func (s *Snowflake) waitNextMillis() int64 {
	currentTime := time.Now().UnixNano() / 1e6
	for currentTime <= s.lastTime {
		currentTime = time.Now().UnixNano() / 1e6
	}
	return currentTime
}
