package main

import "fmt"

func main() {
	var grade float32 = 10.0
	if grade < 3.0 {
		fmt.Printf("Sorry, but you got %.1f, you are reproved.\n", grade)
	} else if grade >= 3.0 && grade < 7.0 {
		fmt.Printf("Ok, you got %.1f, you can still recovery the grade.\n", grade)
	} else {
		fmt.Printf("Nice, you got %.1f, you are approved.\n", grade)
	}
}
