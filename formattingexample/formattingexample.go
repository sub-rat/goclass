package formattingexample

import (
	"errors"
	"fmt"
	"strconv"
)

type point struct {
	x, y int
}

func InitFormattingExample(x int, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, errors.New("x and y cannot be negative")
	}
	p := point{x, y}
	fmt.Printf("format 1: %v \n", p)
	fmt.Printf("format 2: %+v \n", p)
	fmt.Printf("format 3 = %#v \n", p)
	fmt.Printf("Type = %T\n", p)

	fmt.Printf("integer = %d\n", 15)
	fmt.Printf("boolean = %t\n", true)

	fmt.Printf("Binary = %b\n", 8)
	fmt.Printf("Character = %c\n", 65)

	fmt.Printf("Hex = %x\n", 500)

	fmt.Printf("Pointer = %p\n", &p)

	fmt.Printf("format 4 =|%4d|%4d|\n", 1, 2)
	fmt.Printf("format 5 =|%-4d|%-4d|\n", 1, 2)

	fmt.Printf("format 6 =|%-7.2f|%7.2f|\n", 1.2, 1.556)

	fmt.Printf("format 7 = %s\n", "golang")

	fmt.Printf("|%15s|%15s| \n", "---------------", "---------------")
	fmt.Printf("|%15s|%15s| \n", "FirstName", "LastName")
	fmt.Printf("|%15s|%15s| \n", "---------------", "---------------")
	fmt.Printf("|%15s|%15s| \n", "Ram", "Sharma")
	fmt.Printf("|%15s|%15s| \n", "Shyam", "Berma")
	fmt.Printf("|%15s|%15s| \n", "Ram", "Sharma")
	fmt.Printf("|%15s|%15s| \n", "Shyam", "Berma")
	fmt.Printf("|%15s|%15s| \n", "Ram", "Sharma")
	fmt.Printf("|%15s|%15s| \n", "Ram", "Sharma")
	fmt.Printf("|%15s|%15s| \n", "Shyam", "Berma")
	fmt.Printf("|%15s|%15s| \n", "Shyam", "Berma")
	fmt.Printf("|%15s|%15s| \n", "---------------", "---------------")

	return 1, nil
}

func InitNumberParsing(a string) {
	i, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	i, err = strconv.ParseInt("101", 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	i, err = strconv.ParseInt("0x11", 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	f, err := strconv.ParseFloat("10.6", 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)

	j, err := strconv.Atoi("123")
	if err != nil {
		panic(err)
	}
	fmt.Println(j)
}
