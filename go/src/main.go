package main

import (
	"fmt"
	"main/src/sort"
)

func main() {
	arr := []int{2, 4, 3, 2, 6, 6, 3, 2, 10, 4}

	fmt.Println("정렬 전:", arr)

	comparator := func(x, y int) bool {
		return x < y
	}

	sort.MergeSort(arr, comparator)

	fmt.Println("정렬 후:", arr)

}
