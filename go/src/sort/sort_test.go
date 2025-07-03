package sort_test

import (
	"main/src/sort"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
		comp     func(int, int) bool
	}{
		{
			name:     "내림차순 정렬",
			input:    []int{5, 3, 7, 8, 2, 1, 9},
			expected: []int{9, 8, 7, 5, 3, 2, 1},
			comp: func(a, b int) bool {
				return a > b
			},
		},
		{
			name:     "오름차순 정렬",
			input:    []int{5, 3, 7, 8, 2, 1, 9},
			expected: []int{1, 2, 3, 5, 7, 8, 9},
			comp: func(a, b int) bool {
				return a < b
			},
		},
		{
			name:     "정렬된 배열 내림차순 정렬",
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			comp: func(a, b int) bool {
				return a > b
			},
		},
		{
			name:     "정렬된 배열 오름차순 정렬",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			comp: func(a, b int) bool {
				return a < b
			},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				arrToSort := make([]int, len(tc.input))
				copy(arrToSort, tc.input)
				sort.QuickSort(arrToSort, tc.comp)
				if !reflect.DeepEqual(arrToSort, tc.expected) {
					t.Errorf("QuickSort() = %v, want %v", arrToSort, tc.expected)
				}
			})
	}
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
		comp     func(int, int) bool
	}{
		{
			name:     "내림차순 정렬",
			input:    []int{5, 3, 7, 8, 2, 1, 9},
			expected: []int{9, 8, 7, 5, 3, 2, 1},
			comp: func(a, b int) bool {
				return a > b
			},
		},
		{
			name:     "오름차순 정렬",
			input:    []int{5, 3, 7, 8, 2, 1, 9},
			expected: []int{1, 2, 3, 5, 7, 8, 9},
			comp: func(a, b int) bool {
				return a < b
			},
		},
		{
			name:     "정렬된 배열 내림차순 정렬",
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			comp: func(a, b int) bool {
				return a > b
			},
		},
		{
			name:     "정렬된 배열 오름차순 정렬",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			comp: func(a, b int) bool {
				return a < b
			},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				arrToSort := make([]int, len(tc.input))
				copy(arrToSort, tc.input)
				sort.MergeSort(arrToSort, tc.comp)
				if !reflect.DeepEqual(arrToSort, tc.expected) {
					t.Errorf("QuickSort() = %v, want %v", arrToSort, tc.expected)
				}
			})
	}
}
