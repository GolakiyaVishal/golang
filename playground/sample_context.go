package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)
	go worker4(ctx, &wg, ch, "worker4")
	go worker4(ctx, &wg, ch, "worker5")

	go func() {
		i := 1
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- fmt.Sprintf("msg:%d", i):
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

func worker4(ctx context.Context, wg *sync.WaitGroup, ch chan string, name string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(name, msg)
		}
	}
}

func worker5(ctx context.Context, wg *sync.WaitGroup, ch chan string, name string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("%s, %s", name, msg)
		}
	}
}
