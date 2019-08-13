package atomic_test

import (
	"sync/atomic"
	"testing"
	"time"
)

var unid uint64

func TestAddUint64(t *testing.T) {
	//wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		//wg.Add(1)
		go func(i int) {
			id := atomic.AddUint64(&unid, 1)
			println(i, ":", id)
			//wg.Done()
		}(i)
	}
	//wg.Wait()
	time.Sleep(time.Second * 3)
	println("unid:", unid)
}
