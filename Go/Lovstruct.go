package main

import "fmt"

type Record struct {
	First string
	Last  string
	ID    int
}

func main() {
	rec := Record{
		First: "Daisy",
		Last:  "Dharshini",
		ID:    0,
	}
	fmt.Printf("%s %s %d", rec.First, rec.Last, rec.ID)
}
