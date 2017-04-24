package main

import "fmt"

var run string = "bububu" // global??

func main() {
	var a string = "hello"
	var b, c = 1, 2

	fmt.Println(a, b, c)

	d := true

	var e int

	fmt.Println(d, e)

	fmt.Println(run)
}
