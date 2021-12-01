package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2, ok := i.(MyInt)
	if !ok {
		fmt.Println("unexpected type")
		return
	}
	fmt.Println(i2 + 1)
}
