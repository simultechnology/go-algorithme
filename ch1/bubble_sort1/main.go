package main

import (
	"fmt"
	"time"
	"math/rand"
)

func BubbleSort(numbers []int) {
	counter := 0
	k := 0
	flag := true
	for ;flag; {
		flag = false;
		l := len(numbers)
		for i := 0; i < l - 1 - k; i++ {
			counter++
			if numbers[i] > numbers[i + 1] {
				tmp := numbers[i];
				numbers[i] = numbers[i + 1]
				numbers[i + 1] = tmp
				flag = true;
			}
		}
		k++
	}
	fmt.Printf("total : %d\n", counter)
}

func main() {

	var sortNum int
	var randomNumbers[]int

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

	BubbleSort(randomNumbers)

	endTime := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Printf("duration : %d\n", endTime - startTime)

}
