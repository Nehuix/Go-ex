package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}
	v2 := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "green",
		},
		luxury: true,
	}

	fmt.Println(v1)
	fmt.Println(v1.color)
	fmt.Println("--------------------------------")
	fmt.Println(v2)
	fmt.Println(v2.doors)
	fmt.Println(v2.luxury)
}
