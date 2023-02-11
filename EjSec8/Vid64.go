package main

import (
	"fmt"
)

func main() {
	x := 41
	if x == 42 {
		fmt.Println("001")
	} else if x == 41 {
		fmt.Println("es 41")
	} else {
		fmt.Println("no era nada")
	}
}
