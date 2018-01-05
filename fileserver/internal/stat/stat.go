package stat

import (
	"sync"
	"time"
)

// Stat object which collects hits
type Stat struct {
	recordCount int
	sync        sync.Mutex
	lastRate    int
	done        chan struct{}
}

// NewStat returns Stat info
func NewStat() *Stat {
	st := &Stat{
		recordCount: 0,
		lastRate:    0,
	}
	st.rotate()
	return st
}

func (s *Stat) rotate() {
	go func() {
		for {
			tmr := time.NewTimer(time.Second * 1)
			<-tmr.C
			s.lastRate = s.recordCount
			s.reset()
		}
	}()
}

// Hit registers DB hit
func (s *Stat) Hit() {
	s.sync.Lock()
	s.recordCount++
	s.sync.Unlock()
}

func (s *Stat) reset() {
	s.sync.Lock()
	s.recordCount = 0
	s.sync.Unlock()
}

// GetRatePerSecond returns last measured rate per second
func (s *Stat) GetRatePerSecond() int {
	return s.lastRate
}
