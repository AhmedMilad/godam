/*
 * Copyright (c) 2026 Ahmed Milad
 *
 * This Source Code Form is subject to the terms of
 * the Mozilla Public License, v. 2.0.
 * If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

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
