/*
 * Copyright (c) 2020
 * @Author: xiaoweixiang
 */

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := 2
	ans := subarraySum(nums, k)
	fmt.Println("ans:", ans)

}
func subarraySum(nums []int, k int) int {
	ma := make(map[int]int)
	sum := 0
	res := 0
	ma[0] = 1
	for _, num := range nums {
		sum += num
		value, ok := ma[sum-k]
		if ok {
			res += value
		}
		v, okk := ma[sum]
		if okk {
			ma[sum] = v + 1
		} else {
			ma[sum] = 1
		}

		fmt.Println("ma:", ma)
	}
	return res
}
