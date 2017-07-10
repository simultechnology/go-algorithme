package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("start!\n")

	num := 58595869 / 10.0

	ceilNum := int(math.Ceil(num))
	fmt.Printf("%f -> %d\n", num, ceilNum)
}
