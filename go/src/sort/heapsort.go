package sort

import (
	"cmp"
)

func MaxHeapSort[T cmp.Ordered](arr []T) {
	n := len(arr)

	// NOTE. 논리적 트리 생성
	for i := n/2 - 1; i >= 0; i-- {
		heapifyMax(arr, n, i)
	}

	// NOTE. 정렬
	for i := n - 1; i > 0; i-- {
		// NOTE. 맨 앞 데이터는 맨 뒤로 Shift
		arr[0], arr[i] = arr[i], arr[0]
		// NOTE. Shift 데이터를 제외한 나머지 데이터로 논리적 트리 생성
		heapifyMax(arr, i, 0)
	}
}

func heapifyMax[T cmp.Ordered](arr []T, n, i int) {
	largest := i
	left := (2 * i) + 1
	right := (2 * i) + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapifyMax(arr, n, largest)
	}
}

func MinHeapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyMin(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapifyMin(arr, i, 0)
	}
}

func heapifyMin(arr []int, n, i int) {
	smallest := i
	left := (2 * i) + 1
	right := (2 * i) + 2

	if left < n && arr[left] < arr[smallest] {
		smallest = left
	}

	if right < n && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		heapifyMin(arr, n, smallest)
	}
}
