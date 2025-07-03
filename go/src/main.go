package main

import (
	"fmt"
	"main/src/sort"
)

func main() {
	arr := []int{2, 4, 6, 3, 10}

	fmt.Println("정렬 전:", arr)
	// comparator := func(x, y int) bool {
	// 	return x < y
	// }
	sort.MaxHeapSort(arr)

	fmt.Println("정렬 후:", arr)

}
