package main

import (
	"strconv"
	"fmt"
)

func main() {
	f, _ := strconv.ParseFloat("1.23", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123456", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1ece", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("282387", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("124")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
