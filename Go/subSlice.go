package main

import "fmt"

func main() {
	fruits := make([]string, 0)
	fruits = append(fruits, "apple")
	fruits = append(fruits, "orange")
	fruits = append(fruits, "mango")
	californiaFruits := fruits[0:2]
	fmt.Println(fruits[0])
	fmt.Printf("%#v\n", californiaFruits)

}
