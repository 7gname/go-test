package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"context"
)


func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go LaunchProcessor1(ctx)

	<-sigChan
	cancel()
	time.Sleep(time.Second)
}

func LaunchProcessor1(ctx context.Context) {
	fmt.Printf("Start Work\n")
	i := 0
	for ;i < 5;{
		select {
		case <-ctx.Done():
			fmt.Printf("Kill Early\n")
			return
		default:
			fmt.Printf("Doing Work\n")
			time.Sleep(1 * time.Second)
		}
		i++
	}

	fmt.Printf("End Work\n")
}
