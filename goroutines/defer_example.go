package goroutines

import (
	"fmt"

	"github.com/sub-rat/goclass/initials"
)

func InitDeferExample() {
	defer fmt.Println("Program finished 1")
	defer fmt.Println("Program finished 2")
	value := initials.Max(10, 5)
	fmt.Println("Max Value = ", value)
}
