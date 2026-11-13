/*
 * Copyright (c) 2026 Ahmed Milad
 *
 * This Source Code Form is subject to the terms of
 * the Mozilla Public License, v. 2.0.
 * If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

// Package maths provides generic integer math helpers.
package maths

// Integer is a constraint that permits any signed or unsigned integer type,
// including user-defined types whose underlying type is one of them.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Pow returns base raised to exp using binary exponentiation.
// The result may overflow the underlying integer type.
func Pow[T Integer](base T, exp uint64) T {
	res := T(1)

	for exp > 0 {
		if exp&1 == 1 {
			res *= base
		}

		base *= base
		exp >>= 1
	}

	return res
}
