package sync_test

import (
	"sync"
	"testing"
	"reflect"
)

//用 sync.WaitGroup{} 来保证子协程执行完
func TestSyncWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			t.Logf("goroutine %d do...\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//用 sync.Mutex{}对资源加锁，控制协程并发冲突
func TestSyncMutex(t *testing.T) {
	n := 0
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			lock.Lock()
			n = 100-i
			t.Logf("goroutine %d set n = %d", i, n)
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf("last n is %d\n", n)
}

//sync.map
func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			m.Store("aa", i)
			t.Logf("goroutine %d set aa = %d", i, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	v, _ := m.Load("aa")
	t.Logf("last aa is %d", v.(int))
}

func TestSyncPool(t *testing.T) {
	type resource struct {
		Id int
	}
	pool := sync.Pool{}
	for i := 0; i < 10; i ++ {
		pool.Put(resource{i})
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			x := pool.Get()
			xValOf := reflect.ValueOf(x)
			if xValOf.Type().Name() == "resource" {
				println(xValOf.FieldByName("Id").Int())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}