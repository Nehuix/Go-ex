package main

import (
	"fmt"
)

func main() {
	s := struct {
		first   string
		friends map[string]int
		favBeb  []string
	}{
		first: "Nehuen",
		friends: map[string]int{
			"Juli": 123,
			"Mica": 787,
			"Ju":   4544,
		},
		favBeb: []string{
			"Coca",
			"Fanta",
			"Pritty",
		},
	}
	fmt.Println(s)
	fmt.Println("--------------------------------")
	fmt.Println(s.friends)
	fmt.Println(s.first)
	fmt.Println(s.favBeb)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}
}
