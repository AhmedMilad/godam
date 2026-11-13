package algo

import (
	"cmp"
)

func LowerBound[T cmp.Ordered](arr []T, target T) int {

	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func UpperBound[T cmp.Ordered](arr []T, target T) int {

	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
