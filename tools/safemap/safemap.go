package safemap

import "sync"

type SafeMap struct {
	mm   map[interface{}]interface{}
	lock *sync.RWMutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		mm:   make(map[interface{}]interface{}),
		lock: new(sync.RWMutex),
	}
}

func (this *SafeMap) GetLock() *sync.RWMutex {
	return this.lock
}

func (this *SafeMap) GetMap() map[interface{}]interface{} {
	return this.mm
}

func (this *SafeMap) GetMapLen() int {
	return len(this.mm)
}

func (this *SafeMap) Get(key interface{}) (interface{}, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	val, ok := this.mm[key]
	return val, ok
}

func (this *SafeMap) Set(key interface{}, val interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.mm[key] = val
}

func (this *SafeMap) Delete(key interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.mm, key)
}

func (this *SafeMap) Check(key interface{}) bool {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if _, ok := this.mm[key]; ok {
		return true
	}
	return false
}

func (this *SafeMap) ReLoad(m map[interface{}]interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.mm = m
}
