package main
import (
	"fmt"
	"io/ioutil"
)
func main(){
	fileName := "sample_golang.txt"
	//content := []byte("Welcome to ioutil pack file utils")
	content, err := ioutil.ReadFile(fileName)
	// err := ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println("Error found during creation",err)
	} 
	fmt.Println(string(content))

}
