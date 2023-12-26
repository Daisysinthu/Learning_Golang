package main

import "fmt"

func main() {
	var x = "linux"
	switch x {
	case "linux", "bsd", "osx":
		fmt.Println("unix")
	case "windows 10":
		fmt.Println("microsoft os")
	default:
		fmt.Println("Not sure")
	}

}