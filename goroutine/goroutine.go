package goroutine

import (
	"fmt"
	"sync"
)

func TestGoroutine() {

	var wg sync.WaitGroup

	//this states the number of concurrent task to run
	wg.Add(2)

	go func() {
		fmt.Println("working asynchronously")
		wg.Done()
	}()

	go func() {
		fmt.Println("working asynchronously 2")
		wg.Done()
	}()

	fmt.Println("working synchronously")

	//telling the main function to wait for them to ruun b4 it exit
	wg.Wait()
}
