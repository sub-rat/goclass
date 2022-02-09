package examplejson

import (
	"encoding/json"
	"fmt"
)

// type Address struct{
// 	Country string
// 	Street string
// }

type Person struct {
	FirstName string              `json:"first_name"`
	LastName  string              `json:"last_name"`
	Age       int                 `json:"age"`
	Married   bool                `json:"married"`
	Address   map[string]string   `json:"address"`
	Phone     []map[string]string `json:"phone"`
}

type Response struct {
	Page int      `json:"page"`
	Data []Person `json:"data"`
}

func InitExampleJson() {

	// Key , Value
	// Json Object
	// Json Array
	// data type := number, string , bool
	// ex1, _ := json.Marshal(true)
	// fmt.Println(string(ex1))

	// ex2, _ := json.Marshal(1)
	// fmt.Println(string(ex2))

	// ex3, _ := json.Marshal("golang")
	// fmt.Println(string(ex3))

	// exD := []string{"a", "b", "c"}
	// ex4, _ := json.Marshal(exD)
	// fmt.Println(string(ex4))

	// exD2 := map[string]interface{}{"first_name": "Ram", "last_name": "Berma", "age": 24}
	// fmt.Println(exD2)
	// ex5, _ := json.Marshal(exD2)
	// fmt.Println(string(ex5))

	person1 := &Person{
		FirstName: "Ram",
		LastName:  "Berma",
		Age:       24,
		Married:   false,
		Address: map[string]string{
			"country": "Nepal",
			"street":  "buddhanagar",
		},
		Phone: []map[string]string{
			{"home": "98989898"},
			{"office": "897899798"},
		},
	}
	person2 := &Person{
		FirstName: "Ram",
		LastName:  "Berma",
		Age:       24,
		Married:   false,
		Address: map[string]string{
			"country": "Nepal",
			"street":  "buddhanagar",
		},
		Phone: []map[string]string{
			{"home": "98989898"},
			{"office": "897899798"},
		},
	}
	data := []Person{*person1, *person2}
	resp1 := &Response{
		Page: 1,
		Data: data,
	}
	fmt.Printf("resp1=%+v\n", resp1)
	ex6, _ := json.Marshal(resp1)
	fmt.Println(string(ex6))

	exJson := []byte(`{"a":1,"b":"abc","c":[10,11]}`)

	var ex7 map[string]interface{}

	err := json.Unmarshal(exJson, &ex7)
	if err != nil {
		panic(err)
	}
	fmt.Println(ex7)

	// ex2Json := []byte(`{
	// 	"page": 1,
	// 	"data": [
	// 	  {
	// 		"first_name": "Ram1",
	// 		"last_name": "Berma2",
	// 		"age":25
	// 	  },
	// 	  {
	// 		"first_name": "Ram2",
	// 		"last_name": "Berma2",
	// 		"age":16
	// 	  },
	// 	  {
	// 		"first_name": "Ram3",
	// 		"last_name": "Berma3"
	// 	  },
	// 	  {
	// 		"first_name": "Ram4",
	// 		"last_name": "Berma4"
	// 	  }
	// 	]
	//   }`)
	ex2Json := ReadFile("./example_json/person_data.json")

	fmt.Println("===========================")
	fmt.Println("Unmarshaling the Person's record")
	var ex8 Response
	err = json.Unmarshal(ex2Json, &ex8)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", ex8)
	for _, value := range ex8.Data {
		if value.Age > 15 {
			fmt.Printf("%15s %15s\n", value.FirstName, value.LastName)
		}
	}

}
