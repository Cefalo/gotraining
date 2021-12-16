package main

import (
	"fmt"

	"github.com/learning-go-book/package_example/tree/master/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
