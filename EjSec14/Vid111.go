package main

import (
	"fmt"
)

func main() {
	//xi:=[]int{2,3,4,7,11}
	x := sum() //premite pasar el slice,
	//porque la func. solo pide int
	fmt.Println("El total de la fun es: ", x)

}

func sum(x ...int) int {
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
