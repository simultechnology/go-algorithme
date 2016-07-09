package main

import (
	"fmt"
	"math/rand"
	"time"
)

var counter int

func QuickSort(bottom int, top int, data []int) {

	var (
		lower int
		upper int
		div   int
		tmp   int
	)

	if bottom >= top {
		return
	}

	div = data[bottom]

	lower = bottom
	upper = top
	for ; lower < upper; {
		for ; lower <= upper && data[lower] <= div; {
			counter++
			lower++
		}

		for ; lower <= upper && data[upper] > div; {
			counter++
			upper--
		}

		if lower < upper {
			counter++
			tmp = data[lower]
			data[lower] = data[upper]
			data[upper] = tmp
		}
	}

	tmp = data[bottom]
	data[bottom] = data[upper]
	data[upper] = tmp


	QuickSort(bottom, upper - 1, data);
	QuickSort(upper + 1, top, data)

}

func main() {

	var sortNum int
	var randomNumbers []int

	fmt.Printf("ソート件数を入力して下さい。")
	fmt.Scanf("%d", &sortNum)

	//sortNum = 20

	rand.Seed(int64(sortNum))
	//r := rand.New(rand.NewSource(99))

	fmt.Printf("ソート準備\n")
	for i := 0; i < sortNum; i++ {
		randomNumbers = append(randomNumbers, rand.Intn(sortNum))
	}

	fmt.Printf("%v\n", randomNumbers)

	fmt.Printf("ソート開始\n")

	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	QuickSort(0, sortNum-1, randomNumbers)

	endTime := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Printf("total : %d\n", counter)
	fmt.Printf("duration : %d\n", endTime-startTime)
	//fmt.Printf("%v\n", randomNumbers)

}
