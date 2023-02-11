package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("no imprime")
	case (2 == 4):
		fmt.Println("no imprime tamp")
	case (1 == 3):
		fmt.Println("imprime")
		fallthrough
	case (2 == 4):
		fmt.Println("tmb imprime")
		fallthrough
	case (5 == 4):
		fmt.Println("tmb imprime")
		fallthrough
	case (2 == 7):
		fmt.Println("tmb imprime")
	default:
		fmt.Println("cayó default")
	}
}
