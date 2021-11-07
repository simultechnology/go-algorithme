package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/simultechnology/go-algorithme/ch1/sort"
)

func main() {

	var sortNum int
	var randomNumbers []int

	fmt.Printf("ソート件数を入力して下さい。")
	fmt.Scanf("%d", &sortNum)

	rand.Seed(int64(sortNum))
	//r := rand.New(rand.NewSource(99))

	fmt.Printf("ソート準備\n")
	for i := 0; i < sortNum; i++ {
		randomNumbers = append(randomNumbers, rand.Intn(sortNum))
	}

	fmt.Printf("%v\n", randomNumbers)

	fmt.Printf("ソート開始\n")

	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	sort.BubbleSort(randomNumbers)

	fmt.Printf("%v\n", randomNumbers)

	endTime := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Printf("duration : %d\n", endTime-startTime)

}
