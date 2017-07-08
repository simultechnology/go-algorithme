package main

import (
	"./sort"
	"fmt"
	"github.com/simultechnology/go-algorithme/utils"
	"time"
)

func main() {

	var sortNum int
	var randomNumbers []int

	fmt.Printf("ソート件数を入力して下さい。")
	fmt.Scanf("%d", &sortNum)
	randomNumbers = utils.GenerateRandomNumbers(sortNum)
	fmt.Printf("%v\n", randomNumbers)

	fmt.Printf("ソート開始\n")

	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	sort.QuickSort(0, sortNum-1, randomNumbers)

	endTime := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Printf("duration : %d\n", endTime-startTime)
	fmt.Printf("%v\n", randomNumbers)

}
