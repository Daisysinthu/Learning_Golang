package main

import "fmt"

func main() {
	modelToVendor := make(map[string]string)
	modelToVendor["focus"] = "ford"
	modelToVendor["prius"] = "toyato"
	fmt.Printf("%#v\n", modelToVendor["prius"])

}
