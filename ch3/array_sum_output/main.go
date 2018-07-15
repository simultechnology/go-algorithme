package main

import "fmt"

const (
	Nmax = 10
)

func main() {
	fmt.Println("start!")

	var array [Nmax]int
	count := 0
	for {
		fmt.Println("整数を入力して下さい（0を入力すると終了）")
		var inputValue int
		fmt.Scanf("%d", &inputValue)
		if inputValue == 0 {
			break
		} else {
			array[count] = inputValue
			count += 1
			if count >= Nmax {
				break
			}
		}
	}

	sum := 0
	// 合計値を算出
	fmt.Println("--入力されたのは以下の数です--")
	for _, elem := range array {
		fmt.Printf("%d \t", elem)
		sum += elem
	}
	fmt.Printf("\n-----\n以上の数の合計値は%dです。\n", sum)
}

