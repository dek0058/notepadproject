package sort

func MergeSort[T any](arr []T, comparator func(T, T) bool) {
	if len(arr) <= 1 {
		return
	}
	mergeSort(arr, 0, len(arr)-1, comparator)
}

func mergeSort[T any](arr []T, left, right int, comparator func(T, T) bool) {
	if left >= right {
		return
	}
	mid := left + ((right - left) / 2)
	mergeSort(arr, left, mid, comparator)
	mergeSort(arr, mid+1, right, comparator)
	merge(arr, left, mid, right, comparator)
}

func merge[T any](arr []T, left, mid, right int, comparator func(T, T) bool) {
	tmp := make([]T, right-left+1)
	i := left
	j := mid + 1
	k := 0

	for i <= mid && j <= right {
		if comparator(arr[i], arr[j]) {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}

	for idx := i; idx <= mid; idx++ {
		tmp[k] = arr[idx]
		k++
	}

	for idx := j; idx <= right; idx++ {
		tmp[k] = arr[idx]
		k++
	}

	for idx, val := range tmp {
		arr[left+idx] = val
	}
}
