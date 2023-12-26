package main

import "fmt"

type Record struct {
	First string
	Last  string
	ID    int
}

func (r Record) Display() string {
	return fmt.Sprintf("%s,%s,%d", r.First, r.Last, r.ID)
}

func main() {
	rec := Record{
		First: "Daisy",
		Last:  "Dharshini",
		ID:    0,
	}
	fmt.Println(rec.Display())
}
