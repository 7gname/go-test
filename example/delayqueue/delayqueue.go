package delayqueue

import (
	"time"
	"container/list"
)

type Task struct {
	Name     string
	Content  string
	Callback func(s Task)
	DoTime   time.Time
}

var delayQueue []*list.List

const DefaultLengthDelayQueue = 8

func Server(stop <-chan bool) {
	for {
		now := time.Now()
		tasks := delayQueue[now.Unix()%DefaultLengthDelayQueue]
		if tasks.Len() > 0 {
			for e := tasks.Front(); e != nil; e = e.Next() {
				t, ok := e.Value.(Task)
				if ok {
					if t.DoTime.Unix() <= now.Unix() {
						t.Callback(t)
						tasks.Remove(e)
					}
				}else {
					tasks.Remove(e)
				}
			}
		}

		select {
		case c := <-stop:
			if c {
				return
			}
		default:

		}
	}
}

func Add(task Task) error {
	idx := task.DoTime.Unix() % DefaultLengthDelayQueue
	delayQueue[idx].PushBack(task)
	return nil
}

func init() {
	delayQueue = make([]*list.List, DefaultLengthDelayQueue)
	for i, _ := range delayQueue {
		delayQueue[i] = list.New()
	}
}
