package main

import "fmt"

func setZeroes(matrix [][]int) {
	// 定义一个boolean 数组，来定义是否被访问过
	m := len(matrix)
	n := len(matrix[0])
	fmt.Println("m=", m)
	fmt.Println("n=", n)
	var dp [][]bool

	for i := 0; i < m; i++ {
		arr := make([]bool, n)
		dp = append(dp, arr)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Println("i,j=", i, j)
			fmt.Println("value=", matrix[i][j])
			if matrix[i][j] == 0 && !dp[i][j] {
				for k := 0; k < m; k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = 0
						dp[k][j] = true
					}
				}
				for l := 0; l < n; l++ {
					if matrix[i][l] != 0 {
						matrix[i][l] = 0
						dp[i][l] = true
					}
				}
			}
		}
	}
}
