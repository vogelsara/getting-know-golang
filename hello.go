package main

import "fmt"

type book int

var x book
var y int

func main () {
	fmt.Println(x)
	fmt.Printf("%T,\n", x)

	x = 10
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
