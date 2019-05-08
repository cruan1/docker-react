package main

import "fmt"

func main() {

	states := make([]string, 50, 50)

	states = []string{"Alabama", "Alaska", "Arizona"}

	fmt.Println(states)

	for i:=0; i < 3; i++ {

	fmt.Printf("%v %v\n", i, states[i])

}
}
