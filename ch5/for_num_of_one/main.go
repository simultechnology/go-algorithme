package main

import (
	"fmt"
)

func main() {

	var i int

	fmt.Scanf("%d", &i);
	fmt.Printf("the number of 1 in %d is %d\n", i, numOfOne(i))
}

func numOfOne(value int) int {

	var ret int

	for ret = 0; value > 0; value /= 10 {
		if (value % 10 == 1) {
			ret++
		}
	}
	return ret
}