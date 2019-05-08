package main

import (
	"fmt"
)

type vehicle struct {
	doors string
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

	t1 := truck{
		vehicle: vehicle{
			doors: "four",
			color: "green",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: "two",
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(t1.doors,t1.color,t1.fourWheel)

	fmt.Println(s1.doors,s1.color,s1.luxury)

}
