package main

import "fmt"

type Record struct {
	First, Last string
	ID          int
}

func (r Record) String() string {
	return fmt.Sprintf("%s,%s,%d", r.First, r.Last, r.ID)

}

type FileCabinet struct {
	Files  []Record
	nextID int
}

func (f *FileCabinet) AddRecord(first, last string) {
	r := Record{
		First: first,
		Last:  last,
		ID:    f.nextID,
	}
	f.Files = append(f.Files, r)
	f.nextID++

}

func main() {
	fc := &FileCabinet{}
	fc.AddRecord("Rob", "Pike")
	fc.AddRecord("Ken", "Thompson")
	fc.AddRecord("Robert", "Griesemer")
	for _, rec := range fc.Files {
		fmt.Println(rec)
	}
}
