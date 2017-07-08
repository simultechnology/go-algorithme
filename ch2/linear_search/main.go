package main

import (
	"fmt"
	"github.com/simultechnology/go-algorithme/utils"
)

const (
	NOT_FOUND int = -1
	N         int = 20100
)

func LinearSearch(x int, array []int, num int) int {
	fmt.Printf("LinearSearch start!\n")
	var n = 0
	for ; n < num && x != array[n]; n += 1 {
	}

	if n < num {

		return n
	}

	return NOT_FOUND
}
func SentinelLinearSearch(x int, array []int, num int) int {
	fmt.Printf("SentinelLinearSearch start!\n")
	lastValue := array[num-1]
	array[num-1] = x
	var n = 0
	for ; x != array[n]; n += 1 {

	}
	array[num-1] = lastValue
	if n < num-1 {
		return n /* 一番最後以外で一致 */
	}
	if x == lastValue {
		return n
	}
	return NOT_FOUND
}

func main() {

	//var array = make([]int, 0)
	var targetNum int

	fmt.Println("start!")

	/* 適当な配列を作る */
	array := utils.GenerateRandomNumbers(N)

	fmt.Println("何を探しますか？")
	fmt.Scanf("%d", &targetNum)

	result := LinearSearch(targetNum, array[:], N)
	if result == NOT_FOUND {
		fmt.Printf("%d was not found\n", targetNum)
	} else {
		fmt.Printf("%d is %dth\n", targetNum, result)
	}
	result2 := SentinelLinearSearch(targetNum, array[:], N)
	if result2 == NOT_FOUND {
		fmt.Printf("%d was not found\n", targetNum)
	} else {
		fmt.Printf("%d is %dth\n", targetNum, result2)
	}
}
