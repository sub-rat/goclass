package goroutines

import (
	"fmt"
	"time"
)

func Init2() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(8 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
			// default:
			// 	fmt.Println("default")
		}
	}

}
