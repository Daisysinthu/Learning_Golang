package main

import "fmt"

func main() {
	mysli := []string{
		"bob",
		"ann",
		"steve",
	}
	for index, value := range mysli {
		fmt.Printf("%d: %s\n", index, value)
	}

}