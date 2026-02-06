package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	// TASK-2
	// ch := make(chan string, 5)

	// go worker(ch)

	// for msg := range ch {
	// 	fmt.Println(msg)
	// }

	// TASK-1
	// go task("API")
	// go task("DB")
	// go task("Cache")

	// time.Sleep(3 * time.Second)
	// fmt.Print("Done")
}

// TASK-3 main
func main3() {
	wg := sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(3)
	go worket3(&wg, ch, "A")
	go worket3(&wg, ch, "B")
	go worket3(&wg, ch, "C")

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Print("all done")

}

// TASK - 3
func worket3(wg *sync.WaitGroup, ch chan (string), name string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("worder: %s, %d done", name, i)
	}
	wg.Done()
}

// TASK-2
func worker(ch chan (string)) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("worker %d done", i)
	}
	close(ch)
}

// TASK-1
func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s -> %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}
