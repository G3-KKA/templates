package utils

import (
	"sync/atomic"
	"time"
)

// G3KKA Template Library.utils
// Use it as you wish, even if i wrote tests -- you should never use it in production code

// Name speaks for itself
type StopableTicker struct {
	c      chan time.Time
	closed atomic.Bool
}

type StopFunc func()

// Incapsulates channel, provided channel is read-only and cannot be closed by hand
// It WILL be closed inside ticker logic, after calling StopFunc
func (ticker *StopableTicker) Chan() <-chan time.Time {
	return ticker.c
}

// Returns state of the ticker
func (ticker *StopableTicker) Closed() bool {
	return ticker.closed.Load()
}

// Uses time.Ticker inside, adds stopper functionality
// StopFunc closes the channel and stores true in closed flag
// False ticks may occur!
func NewStopableTicker(d time.Duration) (*StopableTicker, StopFunc) {
	stopable := &StopableTicker{
		c:      make(chan time.Time, 1),
		closed: atomic.Bool{},
	}
	go func() {
		ticker := time.NewTicker(d)
		for {
			val := <-ticker.C
			if stopable.closed.Load() {
				close(stopable.c)
				return
			}
			stopable.c <- val

		}
	}()
	stop := func() {
		stopable.closed.Store(true)
	}
	return stopable, stop
}
