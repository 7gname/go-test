package context_study

import (
	"testing"
	"context"
	"fmt"
	"time"
)

func TestContextWithCancel(t *testing.T) {
	over := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, over chan bool) {
		defer func() {
			fmt.Println("defer...")
		}()
		count := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			default:
				fmt.Println("time: ", time.Now().Unix())
			}
			count++
			if count > 10 {
				over <- true
			}
			time.Sleep(time.Second * 1)
		}
	}(ctx, over)

	if <-over {
		fmt.Println("over")
	}
	cancel()
	time.Sleep(time.Second * 3)
}

func TestCtxWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			default:
				fmt.Println("time: ", time.Now().Unix())
			}
			time.Sleep(time.Second * 1)
		}
	}(ctx)

	time.Sleep(time.Second * 10)
	fmt.Println("over")
	cancel()
	time.Sleep(time.Second * 3)
}
