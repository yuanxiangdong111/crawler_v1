package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{4, 5, 6, 7, 1, 2, 3, 0}
	fmt.Println("before = ", nums)
	quick_sort(0, len(nums)-1, nums)
	fmt.Println("after = ", nums)
}

func quick_sort(l, r int, nums []int) {
	if l >= r {
		return
	}

	i, j := l, r
	mid := nums[(l+r)/2]

	for {
		for nums[i] < mid {
			i++
		}
		for nums[j] > mid {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	quick_sort(l, j, nums)
	quick_sort(j+1, r, nums)
}
