/*
 * Copyright (c) 2021
 * @Author: xiaoweixiang
 */

package main

import "math"

func countBalls(lowLimit int, highLimit int) int {
	/**
	 对每位求和，然后存到map里面，统计最多数量的map
	问题1：在go中如何声明hashmap
	问题2：如何遍历
	*/
	m := make(map[int]int)
	value := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		c := i
		for c != 0 {
			sum += c % 10
			c /= 10
		}
		m[sum]++
		value = int(math.Max(float64(m[sum]), float64(value)))
	}
	return value
}
