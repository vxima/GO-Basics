package main

import "fmt"

func main() {
	var v bool
	var f bool
	v = true
	f = false

	fmt.Println(v && f)
	fmt.Println(v || f)
}
