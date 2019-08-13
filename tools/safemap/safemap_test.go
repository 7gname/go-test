package safemap

import (
	"sync"
	"testing"
)

func Test_set(t *testing.T) {
	m := make(map[string]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			m["aa"] = i
			wg.Done()
		}(i)
	}
	wg.Wait()
	v, b := m["aa"]
	if !b {
		t.Fatal("aa not exists")
	}
	println(v)
}
