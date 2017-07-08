package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomNumbers(max int) []int {

	array := make([]int, max)

	for i := 0; i < max; i += 1 {
		rand.Seed(time.Now().UnixNano()) //種にランダムな数字を設定する
		n := rand.Int()
		array[i] = n % max
	}
	return array
}
