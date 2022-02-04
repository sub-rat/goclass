package goroutines

import (
	"fmt"
)

func InitClosingChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			msg, more := <-jobs
			if more {
				fmt.Println("Receving jobs ", msg)
			} else {
				fmt.Println("All Jobs received")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job ", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs")
	<-done
}
