package main

import "strings"

func main() {
	println("start")

	println(mergeAlternately("abc", "pqr"))
	println(mergeAlternately("ab", "pqrs"))
	println(mergeAlternately("abcd", "pq"))

	println("---------------------")

	println(mergeAlternately2("abc", "pqr"))
	println(mergeAlternately2("ab", "pqrs"))
	println(mergeAlternately2("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {

	split1 := strings.Split(word1, "")
	split2 := strings.Split(word2, "")

	mergedArray := make([]string, 0, len(split1)+len(split2))

	for _, w1 := range split1 {
		mergedArray = append(mergedArray, w1)
		if len(split2) > 0 {
			elm := split2[0]
			split2 = split2[1:]
			mergedArray = append(mergedArray, elm)
		}
	}
	if len(split2) > 0 {
		mergedArray = append(mergedArray, split2...)
	}
	return strings.Join(mergedArray, "")
}

func mergeAlternately2(word1 string, word2 string) string {
	maxLength := len(word1) + len(word2)
	mergedArray := make([]string, 0, maxLength)

	i, j := 0, 0
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			mergedArray = append(mergedArray, string(word1[i]))
			i++
		}
		if j < len(word2) {
			mergedArray = append(mergedArray, string(word2[j]))
			j++
		}
	}

	return strings.Join(mergedArray, "")
}
