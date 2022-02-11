package arraypractise

import (
	"errors"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CyclicRotate(A []int, K int) ([]int, error) { //[1] , K=0 []
	// a=[3, 8, 9, 7, 6] ; K = 6

	if K < 0 {
		return nil, errors.New(" K cannot be negative")
	}

	if K == 0 || len(A) == 0 || len(A) == 1 {
		return A, nil
	}

	length := len(A) // length = 5
	if K > length {
		K = 1
		K = K % length // K = 6 % 5= 6/5= 1
	}
	// A[lower:upper] // lower included/ upper not included
	// lhs = A[5-3:5] = A[2:5] = [9, 7, 6]
	lhs := A[length-K : length]
	// fmt.Println("value of lhs=", lhs)
	return append(lhs, A[0:length-K]...), nil
	// A[0:5-3]=A[0:2]=[3, 8]
	//A... = 3,8
	//append([9, 7, 6],3,8) =[9, 7, 6, 3, 8]
}
