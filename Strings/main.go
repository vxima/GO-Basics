package main

import (
	"fmt"
	"strings"
)

func main() {
	//Concat strings
	str := "Hey"
	str2 := "Ho"
	fmt.Println(str + str2)

	//To upper and to lower

	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

	//Replace char
	strEstrela := "****UHU****"
	fmt.Println(strings.ReplaceAll(strEstrela, "*", "!"))

	//TrimSpace
	fmt.Println(strings.TrimSpace(" \t\n Hello, World \n\t\r\n"))

	//Repeat
	fmt.Println(strings.Repeat("Hey!", 3))

}
