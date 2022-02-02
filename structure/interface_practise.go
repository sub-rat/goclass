package structure

import "fmt"

type geometry interface {
	area() float64
	perm() float64
}

func meausure(g geometry) {
	fmt.Println("Area = ", g.area())
	fmt.Println("Perm = ", g.perm())
}
