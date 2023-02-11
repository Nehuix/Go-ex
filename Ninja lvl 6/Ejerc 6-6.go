package main

import (
	"fmt"
)

func main() {
	func() {
		for i:=0; i<11;i++{
			fmt.Println("Es el",i)
		}
	}()
	fmt.Println("Listus")
}

