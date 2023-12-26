package main

import "fmt"

func main() {
	myslice := []string{"apple", "orange", "pear"}
	modelToVendor := map[string]string{
		"focus": "ford",
		"prius": "toyato",
	}
	fmt.Printf("%#v\n", modelToVendor["prius"])
	fmt.Printf("%#v\n", myslice[0])

}