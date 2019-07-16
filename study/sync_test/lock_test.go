package sync

import (
	"sync"
	"testing"
)

func Foo1() {
	i:=0;
	i++
}

var lock sync.Mutex
func Foo2() {
	i:=0
	lock.Lock()
	i++
	lock.Unlock()
}

func BenchmarkFoo1(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Foo1()
	}
}

func BenchmarkFoo2(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Foo2()
	}
}
