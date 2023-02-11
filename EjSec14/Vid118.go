package main

import (
	"fmt"
)

func main() {
	ii := []int{1,2,3,4,5,6,7,8,9,} //la coma final no es obligatoria si está en una linea
									//pero es buena práctica dejarla
	s := sum(ii...)
	fmt.Println("todos",s)

	s2 := even(sum, ii...)
	fmt.Println("pares", s2)
}

func sum(xi ...int) int{  //... porque es uin param variable
	fmt.Printf("%T\n", xi)
	total:=0
	for _, v := range xi{//no necesito el incide poor eso el "_"
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int{
	var yi []int
	for _, v := range vi{
		if v%2 == 0{
			yi = append(yi,v)
		}
	}

	return f(yi...)
}