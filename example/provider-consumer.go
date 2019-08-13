package example

import (
	"fmt"
	"go-test/constant"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	Idx    string `json:"idx"`
	Length int    `json:"length"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func Producer(wg *sync.WaitGroup, ch chan Product, idx int, stop *bool) {
	for !*stop {
		ch <- Product{
			Idx:    string(idx) + "-" + time.Now().Format(constant.TIME_FORMAT),
			Length: rand.Intn(10),
			Width:  rand.Intn(10),
			Height: rand.Intn(10),
		}
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	}
	wg.Done()
}

func Consumer(ch chan Product) {
	for product := range ch {
		fmt.Printf("product idx[%s] length[%d], width[%d], height[%d]\n", product.Idx, product.Length, product.Width, product.Height)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	}
}

func StartProviderConsumer() {
	var wg sync.WaitGroup
	stop := false
	ch := make(chan Product, 10)
	defer close(ch)

	// 创建 5 个生产者
	for i := 0; i < 5; i++ {
		go Producer(&wg, ch, i, &stop)
		wg.Add(1)
	}

	for i := 0; i < 10; i++ {
		go Consumer(ch)
	}

	//stop = true
	wg.Wait()
}
