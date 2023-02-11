package main

import (
	"fmt"
)

// arrays no se usan, se usan slices
func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Money", "Frutilla", "Nuez"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
