/*
 * Copyright (c) 2020
 * @Author: xiaoweixiang
 */

package main

import "math"

func stoneGameII(piles []int) int {
	// 从后往前
	size := len(piles)
	sum := 0
	dp := make([][]int, size)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, size+1)
	}
	for i := size - 1; i >= 0; i-- {
		sum += piles[i]
		for M := 1; M <= size; M++ {
			if i+2*M >= size {
				dp[i][M] = sum
			} else {
				for j := 1; j <= 2*M; j++ {
					dp[i][M] = int(math.Max(float64(dp[i][M]), float64(sum-dp[i+j][int(math.Max(float64(M), float64(j)))])))
				}
			}
		}
	}
	return dp[0][1]
}
