package main

import (
	"fmt"
	"sync"
	"time"
)

type threadSafeSet struct {
	s []interface{}
	sync.RWMutex
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		defer set.RUnlock()
		for elem, value := range set.s {
			ch<-elem
			println("Iter:", elem, value)
		}
		close(ch)
	}()
	return ch
}

func main() {
	th := threadSafeSet{
		s:[]interface{}{"1","2"},
	}
	v := th.Iter()
	time.Sleep(time.Second*5)
	fmt.Printf("%s-%v\n", "ch", v)

}
