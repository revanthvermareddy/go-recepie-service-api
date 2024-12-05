package main

import (
	"fmt"
	"sync"
)

// implementation of a built-in queue
func reader(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		fmt.Printf("Reader: %d received val: %s \n", id, val)
	}
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(4)

	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)
	go reader(5, ch, &wg)
	go reader(6, ch, &wg)

	// pushing the data into the channel
	for i := 0; i < 100; i++ {
		url := fmt.Sprintf("https://example.com/%d", i)
		ch <- url
	}

	close(ch)
	wg.Wait()
}
