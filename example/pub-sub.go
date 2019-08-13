package example

import (
	"errors"
	"sync"
)

type Suber struct {
	Id int
	Ch chan<- string
}

var pub_sub map[string][]Suber
var mu sync.RWMutex

func Pub(name string, content string) error {
	mu.RLock()
	defer mu.RUnlock()
	if subers, ok := pub_sub[name]; !ok || len(subers) == 0 {
		return errors.New("no suber on \"" + name + "\"")
	}
	for _, suber := range pub_sub[name] {
		suber.Ch <- content
	}
	return nil
}

func Sub(name string, suber Suber) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := pub_sub[name]; !ok {
		pub_sub[name] = make([]Suber, 0)
	}
	pub_sub[name] = append(pub_sub[name], suber)
	return nil
}

func init() {
	pub_sub = make(map[string][]Suber)
}
