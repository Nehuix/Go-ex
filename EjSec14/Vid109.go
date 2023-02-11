package main

import (
	"fmt"
)

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := wootang("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func wootang(fn, ls string) (string, bool) {
	a := fmt.Sprint(fn, " ", ls, `, dice "hola"`)
	b := true
	return a, b
}
