package main

import (
	"fmt"
)

func main (){
	for i:=65;i<=90;i++{
		fmt.Println(i)
		for e:=0;e<=2;e++{
			fmt.Printf("\t%#U\n", i)
		}
	}
}
