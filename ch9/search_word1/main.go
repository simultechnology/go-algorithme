package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("start!")
	var targetWord string

	fmt.Printf("検索する言葉を入力して下さい。")
	fmt.Scanf("%s", &targetWord)

	//bytes, _ := ioutil.ReadFile("sample.txt")
	lineNumber, lineContents := SearchInFile("sample.txt", targetWord)
	if lineNumber < 0 {
		fmt.Println("見つかりませんでした")
	} else {
		fmt.Printf("%d : %s\n", lineNumber, lineContents)
	}
	fmt.Println("---------------------------------------------------------")
	result := SimpleSearch("a", "adaptation")
	if result == "" {
		fmt.Println("見つかりませんでした")
	} else {
		fmt.Println("見つかりました")
	}
}
func SearchInFile(fileName string, pattern string) (int, string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for i := 0; ; {
		i += 1
		bytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		line := string(bytes)
		result := SimpleSearch(line, pattern)
		if result != "" {
			fmt.Println("見つかりました")
			return i, line
		}
	}
	return -1, ""
}

func SimpleSearch(text string, pattern string) string {
	for i, _ := range text {
		word := text[i:]
		if len(word) < len(pattern) {
			break
		}
		found := true
		/* パターンの先頭から比較を始める */
		for j := 0; j < len(pattern); j++ {
			if pattern[j] != word[j] {
				found = false
				break /* 一致しなかった */
			}
		}
		if found {
			return text
		}
	}

	return ""
}
