package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main() {

	fmt.Println("start!")

	ok := func(err error) { if err != nil { panic(err) } }

	var nums []int
	for {
		//s := []interface{}{}
		var input string
		_, err := fmt.Scanln(&input)
		ok(err)
		splits := strings.Split(input, ",")
		end := false
		for _, s := range splits {
			if s == "end" {
				end = true
				break
			}
			i, err := strconv.Atoi(s)
			ok(err)
			nums = append(nums, i)
		}
		if end {
			break
		}
	}
	fmt.Printf("input values : %v\n", nums)

	multiGcd := multiGcd(nums[0], nums[1:])
	fmt.Printf("multiGcd : %d\n", multiGcd)

	multiBetterGcd := multiBetterGcd(nums)
	fmt.Printf("multiBetterGcd : %d\n", multiBetterGcd)
}

func gcd(a int, b int) int {
	var i int
	for i = a; i > 0; i-- {
		if a % i == 0 && b % i == 0 {
			break
		}
	}
	return i
}

func multiGcd(a int, b []int) int {
	if len(b) == 1 {
		return gcd(a, b[0])
	}
	return  gcd(a, multiGcd(b[0], b[1:]))
}

func multiBetterGcd(a []int) int {
	if len(a) == 1 {
		fmt.Println("two arguments at least are required!")
		os.Exit(1)
	}
	// 引数が2つの場合、普通にgcdを呼ぶだけ
	if len(a) == 2 {
		return gcd(a[0], a[1])
	}
	return  gcd(a[0], multiBetterGcd(a[1:]))
}
