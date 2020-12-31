/*
 * Copyright (c) 2020
 * @Author: xiaoweixiang
 */

package main

import "fmt"

func main() {
	n, k := 4, 13
	ans := paintingPlan(n, k)
	fmt.Println("ans:", ans)
}
func paintingPlan(n int, k int) int {
	var res int = 0
	mm := map[int]int{}
	if n*n == k {
		return 1
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			tmp := i*n + j*n - i*j
			fmt.Println("tmp:", tmp)
			fmt.Println("mm:", mm)
			if tmp == k {
				x := combine(n, i, mm)
				y := combine(n, j, mm)
				fmt.Println("i:", i)
				fmt.Println("j:", j)
				fmt.Println("x:", x)
				fmt.Println("y:", y)
				res += x * y
			}
		}
	}
	return res
}

func combine(n int, a int, ma map[int]int) int {
	if a == 0 {
		return 1
	}
	m2, ok := ma[a]
	if ok {
		return m2
	}
	k := 1
	for i := n - a + 1; i <= n; i++ {
		k *= i
	}
	m := 1
	for i := 1; i <= a; i++ {
		m *= i
	}
	ma[a] = k / m
	return ma[a]
}
