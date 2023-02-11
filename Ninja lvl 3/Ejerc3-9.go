package main

import (
	"fmt"
)

func main() {
	favSport:= "basket"
	switch favSport {
	case "esqui":
		fmt.Println("Primer")
	case "fulbo":
		fmt.Println("Seg")
	case "hockey":
		fmt.Println("Ter")
	case "basket":
		fmt.Println("Cuar")
	}
	
}