// go run go_pointer.go

package main

import "fmt"

type User struct {
	name string
	age  int
}

func sum(a, b int) *int {
	sum := a + b
	fmt.Println(&sum)
	return &sum
}

func main() {
	var a = 11
	fmt.Println(&a)
	fmt.Println(a)
	fmt.Println("------")

	p := &a
	fmt.Println(&p)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println("------")

	andy := User{
		name: "andy",
		age:  18,
	}

	q := &andy
	fmt.Println((*q).name)
	fmt.Println((*q).age)
	fmt.Println(q.name)
	fmt.Println(q.age)
	fmt.Println("------")

	r := sum(1, 2)
	fmt.Println(&r)
	fmt.Println(r)
	fmt.Println(*r)
}
