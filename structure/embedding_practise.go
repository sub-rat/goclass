package structure

import (
	"fmt"
)

type des interface {
	describe() string
	subscribe() string
}

type base struct {
	des
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base num = %v", b.num)
}

func (b base) subscribe() string {
	return fmt.Sprintf("subscribing %s\n", "")
}

type container struct {
	base
	num int
	str string
}

// Init3 Embedding examples
func Init3() {
	con := container{
		base: base{
			num: 10,
		},
		num: 5,
		str: "test",
	}
	fmt.Printf("container = %+v\n", con)

	fmt.Println(con.num, con.str)
	fmt.Println(con.base.num, con.str)

	fmt.Println(con.describe())
	fmt.Println(con.subscribe())
}

// type B struct {
// 	id         string
// 	created_at time.Time
// 	updated_at time.Time
// 	deleted_at time.Time
// }

// type User struct {
// 	B
// 	name    string
// 	age     int
// 	married bool
// }

// type Admin struct {
// 	B
// 	name     string
// 	perssion []string
// }

// type Customer struct {
// 	B
// 	code string
// }
