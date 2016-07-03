package main

import (
	"fmt"
)

// array
//var mdays [][12]int = [][12]int{
//	{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
//	{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
//}
var mdays [][]int = [][]int{
	{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
}

func isLeap(year int) bool {
	return (year % 4 == 0 && year % 100 != 0 || year % 400 == 0)
}

func dayOfYear(year int, month int, day int) int {

	days := day
	for i := 1; i < month; i++ {
		if isLeap(year){
			fmt.Printf("isLeap!\n")
			days += mdays[1][i - 1]
		} else {
			days += mdays[0][i - 1]
		}
	}
	return days
}

func main() {

	var year, month, day int
	var retryNum int

	fmt.Println("start!")

	for retry := true; retry; {

		fmt.Printf("年 : ")
		fmt.Scanf("%d", &year)
		fmt.Printf("月 : ")
		fmt.Scanf("%d", &month)
		fmt.Printf("日 : ")
		fmt.Scanf("%d", &day)

		fmt.Printf("年内で%d日目です。\n", dayOfYear(year, month, day))

		fmt.Printf("もう一度しますか（1 - はい / 0 - いいえ）")
		fmt.Scanf("%d", &retryNum)

		if retryNum != 1 {
			retry = false
		}
	}

}
