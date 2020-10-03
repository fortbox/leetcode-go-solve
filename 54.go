package awesomeProject
func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0 {
		return nil
	}
	// 定义上、下、左、右，四个变量，分别记录当前到哪里转方向。然后就是写代码了
	// 很简单
	row:=len(matrix)
	col:=len(matrix[0])
	up:=0
	down:=row-1
	left:=0
	right:=col-1
	total:=row*col
	var res[]int
	for total>len(res){
		// 先右，再下，再左，再上
		for i := left; i <=right ; i++ {
			res=append(res,matrix[up][i])
		}
		up++
		for i := up; i <=down ; i++ {
			res=append(res,matrix[i][right])
		}
		right--
		for i := right; i >=left ; i-- {
			res=append(res,matrix[down][i])
		}
		down--
		for i := down; i >=up; i-- {
			res=append(res,matrix[i][left])
		}
		left++
	}
	return res
}