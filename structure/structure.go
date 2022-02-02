package structure

import "fmt"

type person struct {
	name    string
	age     uint8
	address map[string]string
	married bool
}

func Init() {
	p := person{
		name:    "Ram",
		age:     25,
		address: map[string]string{"home": "cde", "office": "abd"},
		married: false,
	}
	fmt.Printf("%+v\n", p)
	p1 := person{
		name:    "Shyam",
		address: map[string]string{"home": "abc", "office": "abd"},
	}

	p1.age = 30
	p1.married = true
	// fmt.Println(p1.name)
	fmt.Printf("%+v\n", p1)

	p2 := newPerson("Hari", 25)
	fmt.Printf("%+v\n", p2)

}

func newPerson(name string, age uint8) *person {
	p := person{name: name, age: age}
	return &p
}
