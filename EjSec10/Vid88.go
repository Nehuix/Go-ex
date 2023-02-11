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

	//v valor, ok: si existe tira true
	v, ok := m["Juan"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Todd"] = 33

	if v, ok := m["James"]; ok {
		fmt.Println("Es el print if", v)

	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

}
