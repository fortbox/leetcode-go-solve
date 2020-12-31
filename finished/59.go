package main

func generateMatrix(n int) [][]int {
	// 定义上、下、左、右，四个变量，分别记录当前到哪里转方向。然后就是写代码了
	// 很简单
	row := n
	col := n
	up := 0
	down := row - 1
	left := 0
	right := col - 1
	total := row * col
	var res [][]int
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}
	index := 1
	for index <= total {
		// 先右，再下，再左，再上
		for i := left; i <= right; i++ {
			res[up][i] = index
			index++
		}
		up++
		for i := up; i <= down; i++ {
			res[i][right] = index
			index++
		}
		right--
		for i := right; i >= left; i-- {
			res[down][i] = index
			index++
		}
		down--
		for i := down; i >= up; i-- {
			res[i][left] = index
			index++
		}
		left++
	}
	return res
}
