package goroutines

import (
	"fmt"
	"time"
)

func Init() {
	routineMain()
	// Channels

	mChannel := make(chan string, 2)

	// mChannel <- "Main Thread"
	go func(msg chan string) {
		// fmt.Println("testing1")
		msg <- "anynomous function goroutine"
	}(mChannel)

	go rountine1(mChannel)

	msg := <-mChannel
	fmt.Println(msg)
	msg = <-mChannel
	fmt.Println(msg)
}

func rountine1(msg chan<- string) {
	// fmt.Println(<-msg)
	msg <- "routine1"
}

func routineMain() {
	fmt.Println("First") // 1 . main thread
	// routine()             //	2. main thread
	fmt.Println("Second") // 3. main thread
	go routine()          // routine thread
	go func() {           // routine thread
		fmt.Println("Third")
	}()

	go func(msg string) { // routine thread
		fmt.Println(msg)
	}("testing")

	time.Sleep(time.Second)
	fmt.Println("Done")
}

func routine() {
	for k := 0; k < 5; k++ {
		fmt.Println(k)
	}
}

// func routine() {
// 	for i := 0; i < 10000; i++ {
// 		for j := 0; j < 10000; j++ {
// 			for k := 0; k < 10000; k++ {
// 				fmt.Println(i, j, k)
// 			}
// 		}
// 	}
// }
