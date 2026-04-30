package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c, d = 1, 2, 3, 4
	fmt.Printf("a + b + c + d =%d", a+b+c+d)
	open, _ := os.Open("readme.txt")
	fmt.Printf("open =%v", open)
}
