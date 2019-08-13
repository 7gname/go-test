package counter

import "sync/atomic"

type Counter struct {
	C uint64
}

func NewCounter() Counter {
	return Counter{C: 0}
}

func (c *Counter) Add() uint64 {
	return atomic.AddUint64(&c.C, 1)
}

func (c *Counter) Get() uint64 {
	return atomic.LoadUint64(&c.C)
}
