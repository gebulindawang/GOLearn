package main

import (
	"fmt"
	"slices"
)

func main() {
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	slice := slices.Clone(arr1[:])

	fmt.Println(slice)
}
