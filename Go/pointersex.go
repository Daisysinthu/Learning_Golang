package main

import "fmt"

func changeStr(str *string) {
	*str = "Siddhu"
	fmt.Println("Inside function", *str)

}

func main() {
	str := "Sindhu"
	fmt.Println("In main func", str)
	changeStr(&str)
	fmt.Println("In main func", str)

}