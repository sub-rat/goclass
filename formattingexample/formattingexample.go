package formattingexample

import "fmt"

type point struct {
	x, y int
}

func InitFormattingExample() {
	p := point{10, 20}
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

}
