package delayqueue

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	cb := func(v Task) {
		fmt.Printf("now:%d\ttash:%s\tdelay at:%d\n", time.Now().Unix(), v.Name, v.DoTime.Unix())
	}
	stop := make(chan bool)
	go Server(stop)

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second)
			now := time.Now()
			Add(Task{
				Name:     fmt.Sprintf("task-%03d", i),
				Content:  fmt.Sprintf("%03d-%d", i, now.Unix()),
				Callback: cb,
				DoTime:   now.Add(time.Second * time.Duration(rand.Intn(30))),
			})
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
