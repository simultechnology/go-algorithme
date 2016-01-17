package main

import "fmt"

var STATIONS = []string{
	"横浜", "武蔵小杉", "品川", "渋谷", "新橋", "溜池山王",
}

type AdjacencyMatrix struct {
	name      string
	distances []int
}

func main() {

	var adjancencyMatrixs []*AdjacencyMatrix

	yokohama := &AdjacencyMatrix{name: "横浜", distances: []int{0, 12, 28, 0, 0, 0}}
	adjancencyMatrixs = append(adjancencyMatrixs, yokohama)
	musashiKosugi := &AdjacencyMatrix{name: "武蔵小杉", distances: []int{12, 0, 10, 13, 0, 0}}
	adjancencyMatrixs = append(adjancencyMatrixs, musashiKosugi)
	shinagawa := &AdjacencyMatrix{name: "品川", distances: []int{28, 10, 0, 11, 7, 0}}
	adjancencyMatrixs = append(adjancencyMatrixs, shinagawa)
	shibuya := &AdjacencyMatrix{name: "渋谷", distances: []int{0, 13, 11, 0, 0, 9}}
	adjancencyMatrixs = append(adjancencyMatrixs, shibuya)
	shinbashi := &AdjacencyMatrix{name: "新橋", distances: []int{0, 0, 7, 0, 0, 4}}
	adjancencyMatrixs = append(adjancencyMatrixs, shinbashi)
	tameikeSanno := &AdjacencyMatrix{name: "溜池山王", distances: []int{0, 0, 0, 9, 4, 0}}
	adjancencyMatrixs = append(adjancencyMatrixs, tameikeSanno)

	for _, a := range adjancencyMatrixs {
		fmt.Printf("%s : ", a.name)
		for i, d := range a.distances {
			if d > 0 {
				fmt.Printf("->%s(%d分) ", STATIONS[i], d)
			}
		}
		fmt.Println("")
	}

}
