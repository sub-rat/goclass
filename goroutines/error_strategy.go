package goroutines

import (
	"errors"
	"fmt"
)

func InitErrorStrategy() {

	value, err := f1(5)
	if err != nil {
		panic(err)
	}
	fmt.Println("Value = ", value)
}

func f1(arg int) (int, error) {
	if arg < 0 {
		return 0, errors.New(" Cant work with negative numbers")
	}
	return arg + 1, nil
}
