/*
 * Copyright (c) 2020
 * @Author: xiaoweixiang
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	sweetness := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 5
	ans := maximizeSweetness(sweetness, k)
	fmt.Println("ans:", ans)
}
func maximizeSweetness(sweetness []int, K int) int {
	// 很巧的解法，用二分，参考了题解思路，二分竟然可以这么用
	left, right := math.MaxInt64, 0
	for _, i := range sweetness {
		right += i
		left = int(math.Min(float64(left), float64(i)))
	}
	for left < right {
		mid := left + (right-left+1)/2
		fmt.Println("mid:", mid)
		if canCut(sweetness, mid, K) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func canCut(sweetness []int, mid int, k int) bool {
	sum, count := 0, 0
	for _, i := range sweetness {
		fmt.Println("i:", i)
		sum += i
		if sum >= mid {
			fmt.Println("sum:", sum, ",mid:", mid, ",i:", i)
			count++
			sum = 0
		}
	}
	fmt.Println("count:", count)
	res := count >= k+1
	fmt.Println("res:", res)
	return res
}
