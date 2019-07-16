package example

import (
	"testing"
	"time"
)

func TestPub_Sub(t *testing.T)  {
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch := make(chan string)
			Sub("foo", Suber{i, ch})
			for {
				select {
				case content := <-ch:
					t.Log(i, ":", content)
				}
			}
		}(i)
	}
	go func() {
		for {
			time.Sleep(time.Second * 3)
			Pub("foo", time.Now().String())
		}
	}()
}
