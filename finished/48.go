package finished

func rotate(matrix [][]int) {
	var n = len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
