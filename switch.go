package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2;
	fmt.Println("write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	//var t time.Time = time.Now() // full declare
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	var whatAmI func(i interface{}) = func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}

	whatAmI("a")
	whatAmI(true)
	whatAmI(1)
}
