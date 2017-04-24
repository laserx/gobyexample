package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "alice", age: 19})
	fmt.Println(person{name: "fred"})
	fmt.Println(&person{name: "ann", age:15})
	s := person{name: "Sean", age:30}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s)
}
