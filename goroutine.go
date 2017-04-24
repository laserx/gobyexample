package main

import "fmt"

func f(form string) {
	for i := 0; i < 3; i++ {
		fmt.Println(form, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
