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

	//m["Todd"]=33

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "James")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)

}
