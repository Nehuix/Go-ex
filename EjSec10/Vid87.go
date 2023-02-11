package main

import (
	"fmt"
)

// arrays no se usan, se usan slices
func main() {
	//map [key]value
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	v, ok := m["Juan"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println("Es el print if", v)

	}

}
