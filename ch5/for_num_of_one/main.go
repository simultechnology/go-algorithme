package main

import (
	"fmt"
)

func main() {

	var i int

	fmt.Scanf("%d", &i);
	fmt.Printf("the number of 1 in %d is %d\n", i, numOfOne(i))
	fmt.Printf("the number of 1 in %d is %d\n", i, numOfOneRecursive(i))
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

func numOfOneRecursive(value int) int {

	ret := 0

	if value == 0 {
		return ret
	}
	if value % 10 == 1 { /* 一番下の位が1 */
		ret = 1
	}

	/* 10で割って桁を一つずらし、再びnum_of_oneで調べる */
	return ret + numOfOneRecursive(value / 10)
}