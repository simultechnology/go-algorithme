package main

import "fmt"

func main() {
	fmt.Println("start!")
	slice1 := make([]string, 20)

	slice1 = append(slice1, "a")
	slice1 = append(slice1, "b")
	slice1 = append(slice1, "c")
	fmt.Printf("%v", slice1)

	slice2 := make([]string, 0, 20)
	fmt.Printf("%v", slice2)

	slice2 = append(slice2, "a")
	slice2 = append(slice2, "b")
	slice2 = append(slice2, "c")
	fmt.Printf("%v", slice2)

	slice3 := make([]string, 0)
	fmt.Printf("%v", slice3)
}
