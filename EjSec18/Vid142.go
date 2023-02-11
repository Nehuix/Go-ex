package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByBame []Person

func (a ByBame) Len() int           { return len(a) }
func (a ByBame) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBame) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)

	fmt.Println(people)
	sort.Sort(ByBame(people))
	fmt.Println(people)
}
