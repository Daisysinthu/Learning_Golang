package main

import "fmt"

func main() {
	mysli := map[int]string{
		1: "bob",
		2: "ann",
		3: "steve",
	}
	for key, value := range mysli {
		fmt.Printf("%d: %s\n", key, value)
	}

}