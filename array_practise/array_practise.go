package arraypractise

import "fmt"

func Init() {
	arrayPractise()
	slicesPractise()
}

func slicesPractise() {
	s := make([]int, 2, 3)
	fmt.Println("Slice inital = ", s)
	fmt.Printf("length = %d, capacity = %d \n", len(s), cap(s))
	s[0] = 1
	s[1] = 2
	fmt.Println("Slice 1 = ", s)

	s = append(s, 4)
	s = append(s, 3)
	fmt.Println("Appended Slice", s)
	fmt.Printf("length = %d, capacity = %d \n", len(s), cap(s))

	s1 := make([]int, len(s))
	copy(s1, s)
	fmt.Println("S1 after copy", s1)

	l := s[2:4] // s[ lower_bound : upper_bound ] lower is included; upper is excluded
	fmt.Println("slice 2 =", l)

	l = s[:3] // 0 to 3-1(2)
	fmt.Println("slice 3 =", l)

	l = s[1:] // 1 to last index
	fmt.Println("slice 4 =", l)

	singlLineSlice := []string{"a", "b", "c"}
	fmt.Println("Slice 5 =", singlLineSlice)

	tD := make([][]int, 4)
	for i := 0; i < 4; i++ { // i = 1
		internaArrayLen := i + 1 // 1+1 =2
		tD[i] = make([]int, internaArrayLen)
		for j := 0; j < internaArrayLen; j++ { //j =0; j= 1
			tD[i][j] = i + j
		}
	}
	fmt.Println("2d Slice = ", tD)

}

func arrayPractise() {
	var a [4]int // [0, 0, 0, 0] // variable [length_of_array]data_type
	fmt.Println("array inital =", a)
	a[2] = 4 // a[index] = value
	fmt.Println("Array 1 =", a)
	fmt.Println("2nd index = ", a[2])

	fmt.Println("length of array", len(a))

	b := [3]string{"a", "b", "c"} // initialize array in one line
	fmt.Println("Array 2 =", b)

	var tD [2][3]int         // [[0,0,0][0,0,0]]
	for i := 0; i < 2; i++ { // i =0; i=1
		for j := 0; j < 3; j++ { // j = 0; j=1; j=2 ;;  j = 0; j=1; j=2
			tD[i][j] = i + j //tD[0][0] = 0; tD[0][1] = 1
			// 00, 01, 02, 10, 11, 12
		}
	}
	// [[size=3],[size =3]] = size 2
	fmt.Println("2d array = ", tD) // [[0 1 2] [1 2 3]]
}
