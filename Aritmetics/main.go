package main

import "fmt"

func main() {
	//declaraçao longa
	var a int
	var b int
	a = 2
	b = 3
	fmt.Println(a + b)

	//declaraçao rapida, ja inferindo o tipo

	x := 5
	y := 5
	fmt.Println(x * y)
}
