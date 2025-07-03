package sort

func QuickSort[T any](arr []T, comparator func(T, T) bool) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1, comparator)
}

func quickSort[T any](arr []T, low, high int, comparator func(T, T) bool) {
	if low >= high {
		return
	}

	pivot := partition(arr, low, high, comparator)
	quickSort(arr, low, pivot-1, comparator)
	quickSort(arr, pivot+1, high, comparator)
}

func partition[T any](arr []T, low, high int, comparator func(T, T) bool) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if comparator(arr[j], pivot) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
