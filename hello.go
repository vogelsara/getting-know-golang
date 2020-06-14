package main

import "fmt"

type book int

var x book

func main () {
	x = 42
	fmt.Println(x)
	fmt.Printf("%T,\n", x)
}
