package main
import "fmt"

const STATION_NUMBER = 6

var STATIONS = []string {
	"鎌倉", "藤沢", "横浜", "横須賀", "茅ヶ崎", "東京",
}

var adjacencyMatrix = [][]int {
	{0, 1, 1, 1, 0, 0},
	{1, 0, 1, 0, 1, 0},
	{1, 1, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0},
}

func main() {
	for i := 0; i < STATION_NUMBER; i++ {
		fmt.Printf("%s : ", STATIONS[i])
		for j := 0; j < STATION_NUMBER; j++ {
			if adjacencyMatrix[i][j] == 1 {
				fmt.Printf("->%s ", STATIONS[j])
			}
		}
		fmt.Println("")
	}
}
