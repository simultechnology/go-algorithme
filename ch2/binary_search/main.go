package main

import (
	"fmt"
	"github.com/simultechnology/go-algorithme/ch1/sort"
	"github.com/simultechnology/go-algorithme/utils"
	"math"
	"math/rand"
	"time"
)

type POSTION int

const (
	FIRST POSTION = iota
	LAST
)

const (
	NOT_FOUND int = -1
	N             = 20
)

func BinarySearch(x int, array []int, left int, right int, positon POSTION) int {
	var mid int
	for left < right {
		fmt.Printf("left : %d\n", left)
		fmt.Printf("right : %d\n", right)

		switch positon {

		case FIRST:
			fmt.Printf("case FIRST start!\n")
			//対象の値が見つかったらすぐ返却する
			//if array[mid] == x {
			//	return mid
			//}
			mid = (left + right) / 2

			if array[mid] < x {
				left = mid + 1
				//left = mid
			} else {
				//right = mid - 1
				right = mid
			}
			//同じ数が配列に登録されている場合、配列の先頭を返すようにする
			if array[left] == x {
				return left
			}

		case LAST:
			fmt.Printf("case LAST start!\n")
			mid = int(math.Ceil(float64(left+right) / 2))
			if array[mid] > x {
				right = mid - 1
				//left = mid
			} else {
				//right = mid - 1
				left = mid
			}
			//同じ数が配列に登録されている場合、配列の最後を返すようにする
			if array[right] == x {
				return right
			}
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
	fmt.Printf("generated Array is %v\n", array)

	fmt.Println("何を探しますか？")
	var targetNum int

	fmt.Scanf("%d", &targetNum)

	var pos POSTION
	rand.Seed(time.Now().UnixNano()) //種にランダムな数字を設定する
	n := rand.Int()
	if n%2 == 0 {
		pos = FIRST
	} else {
		pos = LAST
	}
	fmt.Printf("%v\n", pos)

	result := BinarySearch(targetNum, array, 0, N-1, LAST)
	//result := BinarySearch(targetNum, []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 3, 3}, 0, N-1, pos)
	if result == NOT_FOUND {
		fmt.Printf("%d was not found\n", targetNum)
	} else {
		fmt.Printf("%d is %dth\n", targetNum, result)
	}
}
