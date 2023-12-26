package main

import "fmt"

func divide(x, y int) (int, int) {
	result := x / y
	reminder := x % y
	return result, reminder
}

func main() {
	a, m := divide(3, 3)
	fmt.Println(a, m)
}
