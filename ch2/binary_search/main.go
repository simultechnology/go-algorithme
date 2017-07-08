package main

import (
	"github.com/golang/go/src/fmt"
	"github.com/simultechnology/go-algorithme/ch1/sort"
	"github.com/simultechnology/go-algorithme/utils"
)

const (
	NOT_FOUND int = -1
	N             = 200000
)

func BinarySearch(x int, array []int, left int, right int) int {
	var mid int
	for left < right {
		mid = (left + right) / 2
		if array[mid] == x {
			return mid
		}
		if array[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return NOT_FOUND
}
func main() {
	fmt.Println("start!")
	//ランダムな配列を作成する
	array := utils.GenerateRandomNumbers(N)
	//ソートする
	//sort.BubbleSort(array)
	sort.QuickSort(0, N-1, array)

	fmt.Println("何を探しますか？")
	var targetNum int

	fmt.Scanf("%d", &targetNum)

	result := BinarySearch(targetNum, array, 0, N-1)
	if result == NOT_FOUND {
		fmt.Printf("%d was not found\n", targetNum)
	} else {
		fmt.Printf("%d is %dth\n", targetNum, result)
	}
}
