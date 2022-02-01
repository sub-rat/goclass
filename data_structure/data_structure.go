package datastructure

import "fmt"

func Init() {
	m := make(map[string]string) // map[key-type]val-type

	m["nepal"] = "kathmandu"
	m["b"] = "Nepal"
	m["a"] = "Pokhara"
	m["c"] = "Bandipur"
	delete(m, "c")
	practiseMap(m)

	practiseRange()
	practiseVariadicFunction()

	// annynomous functions
	sum := func(nums ...int) {
		total := 0
		for _, num := range nums {
			total += num
		}
		fmt.Println(total)
	}
	sum(1, 2, 3)

	// clousers
	nextSeq := seq()
	fmt.Println(nextSeq()) // i++ = 1
	fmt.Println(nextSeq()) // i++ = 2
	fmt.Println(nextSeq())

	nextSeq1 := seq()
	fmt.Println(nextSeq1()) // i++ = 1
	fmt.Println(nextSeq1()) // i++ = 2
}

//closure
func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func practiseVariadicFunction() { // n no of arguments of type int
	nums := []int{6, 2, 1, 4, 8, 3}
	sum(6, 2, 1, 4, 8, 3)
	sum(1, 2)
	sum(1, 2, 3, 7, 4)
	sum(nums...)
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func practiseRange() {
	arr1 := []int{6, 2, 1, 4, 8, 3}
	for index := 0; index < len(arr1); index++ {
		fmt.Println(index)
		fmt.Println(arr1[index])
	}

	for _, value := range arr1 { // iterates over the array arr1
		// fmt.Println(index)
		fmt.Println(value)
	}

	// inline initialization of map
	m2 := map[string]string{"a": "kathmandu", "b": "Pokhara"}

	for k, v := range m2 { // iterates over the map m2
		fmt.Println("key = ", k)
		fmt.Println("value = ", v)
	}

}

func practiseMap(country map[string]string) {
	// key is always unique
	fmt.Println("len = ", len(country))

	if value, ok := country["nepal"]; ok { // ok = true if key exists  in map
		fmt.Println(ok)
		fmt.Println(value)
	}

	// m2 := map[string]string{"a": "kathmandu", "b": "Pokhara"}
	// fmt.Println("m2", m2)
}
