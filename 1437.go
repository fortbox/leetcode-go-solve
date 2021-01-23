/*
 * Copyright (c) 2021
 * @Author: xiaoweixiang
 */

package leetcode_go_solve

func kLengthApart(nums []int, k int) bool {
	s := -k - 1
	for i, v := range nums {
		if v == 1 {
			if i-s <= k {
				return false
			}
			s = i
		}
	}
	return true
}
