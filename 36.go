package awesomeProject

func isValidSudoku(board[][]byte)bool{
	var row,col,sbox [9][9] bool
	for i:=0;i<9;i++ {
		for j := 0; j < 9; j++ {
			if board[i][j]!='.' {
				num:=board[i][j]-'1'
				index:=(i/3)*3+j/3
				if row[i][num]==true {
					return false
				}else {
					row[i][num]=true
				}
				if col[j][num]==true {
					return false
				}else {
					col[j][num]=true
				}
				if sbox[index][num]==true {
					return false
				}else {
					sbox[index][num]=true
				}
			}
		}
	}
	return true
}
