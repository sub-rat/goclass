package structure

import (
	"fmt"
	"math"
)

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perm() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perm() float64 {
	return 2 * math.Pi * c.radius
}

// Init2 x width of rect, y height of rect, z radius of circe
func Init2(x, y, z float64) {
	r := rect{width: x, height: y}
	c := circle{radius: z}
	// fmt.Println("Area = ", r.width*r.height)

	// fmt.Println("Area of rect = ", r.area())
	// fmt.Println("Perm of rect = ", r.perm())

	fmt.Println("Rect")
	meausure(r)
	fmt.Println("Circle")
	meausure(c)

	// r2 := &r
	// fmt.Println("Area of rect = ", r2.area())
	// fmt.Println("Perm of rect = ", r2.perm())

	// fmt.Println("Area of circle = ", c.area())
	// fmt.Println("Perm of circle = ", c.perm())

}
