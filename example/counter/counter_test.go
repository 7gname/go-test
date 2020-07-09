package counter

import (
	"math/rand"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	c := NewCounter()
	for i := 0; i < 100; i++ {
		go func() {
			for {
				time.Sleep(time.Second * time.Duration(rand.Intn(3)))
				println(c.Add())
			}
		}()
	}
}
