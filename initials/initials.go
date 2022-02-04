package initials

import (
	"errors"
	"fmt"
)

func Init() {
	fmt.Println("tests")
	// var a, b, c = 3, 4, "foo"
	// var a int = 1
	// var p *int
	// p = &a
	// c := "a"
	// // var d float32 = 1.1
	// // a := 2      // a = a + 2
	// b := 3 << 1 // 0110
	// if b < 5 {
	// 	fmt.Println(p)
	// 	fmt.Println("value of b = ", b)
	// 	fmt.Printf("%T\n", c)
	// }
	// ifElseCheck(5)
	// conditionalStatementExample()
	// add(a, d)
	// var greeting = "Hello world!"
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.ToLower(greeting))
	// fmt.Println(strings.Count(greeting, "H"))

	//Write a program using loop up n number
	//and then prints the sum of the even and odd integers at the end?

	even, odd, err := SumOfEvenOdd(5, 10)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
	fmt.Println("even: ", even)
	fmt.Println("odd: ", odd)

	even, odd, err = SumOfEvenOdd(100, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("even: ", even)
	fmt.Println("odd: ", odd)

	// even, odd, err = sumOfEvenOdd(200, 100)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("even: ", even)
	// fmt.Println("odd: ", odd)

	maximumValue := Max(5, 10)
	fmt.Println("Maximun no is = ", maximumValue)

	// a := 5
	// b := 10
	// fmt.Println("Before Swap , value of a = ", a)
	// fmt.Println("Before Swap , value of b = ", b)
	// swap(&a, &b) // & is used to pass the address of a and b
	// fmt.Println("After Swap , value of a = ", a)
	// fmt.Println("After Swap , value of b = ", b)
}

func Swap(a, b *int) {
	var temp int
	temp = *a // *a means value of a ie. value of a = 5 ; temp = 5
	*a = *b
	*b = temp
	fmt.Println("In Swap func value of a = ", *a)
	fmt.Println("In Swap func value of b = ", *b)
}

func SumOfEvenOdd(n, start int) (int, int, error) {
	even := 0
	odd := 0
	if start >= n {
		return 0, 0, errors.New("start is greater than n")
	}
	for i := start; i <= n; i++ {
		if i%2 == 0 {
			even = even + i
		} else {
			odd = odd + i
		}
	}
	return even, odd, nil
}

func Max(num1, num2 int) int {
	defer fmt.Println("End of the Max function")
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

// func conditionalStatementExample() {
// 	sum := 0
// 	for k := 0; k < 2; k++ {

// 	out:
// 		for j := 0; j < 3; j++ {
// 			// temp := false
// 			for i := 0; i < 10; i++ {
// 				if i == 7 {
// 					// temp = true
// 					break out
// 				}
// 				if i == 5 {
// 					continue
// 				}
// 				sum += i
// 				fmt.Println(sum)
// 			}
// 			// if temp {
// 			// 	break
// 			// }
// 		}
// 	}
// 	fmt.Println("Total sum = ", sum)

// 	sum = 1
// 	for sum < 1000 {
// 		sum += sum // sum = sum + sum
// 	}
// 	fmt.Println("Total sum without initial and post statement = ", sum)

// 	// for {
// 	// 	// forever loop
// 	// }

// }

// func ifElseCheck(x int) {
// 	// x := 1 + 3
// 	// if x > 9 {
// 	// 	fmt.Println(x)
// 	// } else if x > 7 {
// 	// 	fmt.Println("on else if 1= ", x)
// 	// } else if x > 5 {
// 	// 	fmt.Println("on else if 2= ", x)
// 	// } else {
// 	// 	fmt.Println("on else = ", x)
// 	// }

// 	switch x {
// 	case 9:
// 		fmt.Println(x)
// 	case 7:
// 		fmt.Println("on else if 1= ", x)
// 	case 5:
// 		fmt.Println("on else if 2= ", x)
// 	default:
// 		fmt.Println("on else = ", x)
// 	}

// }

// // func add(a int, b float32) {

// // 	c := float32(a) + b
// // 	c -= b // c = c - b
// // 	fmt.Println(c)
// // }
