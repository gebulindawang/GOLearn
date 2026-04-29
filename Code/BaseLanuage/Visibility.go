package main

import (
	"fmt"
)

func main() {
	a, b := getNum(1, 2)
	fmt.Printf("a + b =%d,a - b =%d", a, b)
}

func getNum(x, y int) (a, b int) {

	return x + y, x - y
}
