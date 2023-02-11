package main

import (
	"fmt"
)

func main() {
	vari := foo(2, 3, 4, 7, 11)
	fmt.Println("El total de la fun es: ", vari)

}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("Para el indice en,", i, " le sumamos, ", v, " al total y da", sum)

	}
	fmt.Println("El total es: ", sum)
	return sum
}
