package main

import (
	"fmt"
)

func main() {
	foo()

	func() {
		fmt.Println("Func anonima corrió")
	}()

	func(x int) {
		fmt.Println("Numero", x)
	}(42)
}

func foo() {
	fmt.Println("hello from foo")
}
