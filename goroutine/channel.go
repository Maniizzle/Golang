package goroutine

import (
	"fmt"
	"sync"
)

func ChannelTest() {

	var wq sync.WaitGroup
	ch := make(chan string)
	wq.Add(1)

	go func() {
		ch <- "the message"
	}()

	go func() {
		fmt.Println(<-ch)
		wq.Done()
	}()

	wq.Wait()
}

func SelectWithChannel() {

	//this functions leads to a deadlock and eventually panic because there is no message in any channel
	ch1, ch2 := make(chan string), make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case msg2 := <-ch2:
		fmt.Println(msg2)

	}

}
func SelectWithChannelOne() {

	//only default will run because no message in the channel to activate a case
	ch1, ch2 := make(chan string), make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case msg2 := <-ch2:
		fmt.Println(msg2)

	default:
		fmt.Println("No message in channel")

	}

}

func SelectWithTwoChannels() {

	// RUNTIME WILL PICK ANY SELECT CASE

	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "FIRST MESSAGE"
	}()

	go func() {
		ch2 <- "SECOND MESSAGE"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case msg2 := <-ch2:
		fmt.Println(msg2)

	}

}

func SelectWithTwoChannelsExample3() {

	// RUNTIME WILL PICK ANY SELECT CASE

	ch1 := make(chan string)

	go func() {
		ch1 <- "FIRST MESSAGE"
		ch1 <- "SECOND MESSAGE"
		close(ch1)
	}()

	for msg := range ch1 {
		fmt.Println(msg)
	}

}

func LoopingOverChannel() {
	channelOne := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			channelOne <- i
		}
		close(channelOne)
	}()

	for msg := range channelOne {
		fmt.Println(msg)
	}
}
