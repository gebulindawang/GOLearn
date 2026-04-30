package main

import "fmt"

type session uint8

const (
	Spring session = iota
	Summer
	Autumn
	Winter
)

func main() {
	fmt.Printf("Spring = %d,Summer = %d,Autumn = %d,Winter = %d", Spring, Summer, Autumn, Winter)
}
