package main

import (
	"fmt"
)

func main() {
	switch "Holi"{
	case "asd":
		fmt.Println("no imprime")
	case "qwe":
		fmt.Println("no imprime tamp")
	case "Hi":
		fmt.Println("holissss")
		
	case "(2 == 4)", "holi":
		fmt.Println("tmb imprime")
		
	case "(5 == 4)", "Holi":
		fmt.Println("5=4")
		
	case "hol":
		fmt.Println("tmb imprime")
	default:
		fmt.Println("cayó default")
	}
}
