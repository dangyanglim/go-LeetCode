package main

func isValidSudoku(board [][]byte) bool {
	var t map[byte]int
	t = make(map[byte]int)
	for i := 0; i < len(board); i++ {
		t = make(map[byte]int)
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				_, ok := t[board[i][j]]
				if ok {
					return false
				} else {
					t[board[i][j]]++
				}
			}
		}
	}
	for i := 0; i < len(board[0]); i++ {
		t = make(map[byte]int)
		for j := 0; j < len(board); j++ {
			if board[j][i] != '.' {
				_, ok := t[board[j][i]]
				if ok {
					return false
				} else {
					t[board[j][i]]++
				}
			}
		}
	}
	jStart := 0
	iStart := 0
	for l := 0; l < 3; l++ {
		for k := 0; k < 3; k++ {
			t = make(map[byte]int)
			for i := iStart; i < iStart+3; i++ {
				for j := jStart; j < jStart+3; j++ {
					if board[j][i] != '.' {
						_, ok := t[board[j][i]]
						if ok {
							return false
						} else {
							t[board[j][i]]++
						}
					}
				}
			}
			iStart += 3
		}
		jStart += 3
		iStart = 0
	}

	return true
}

func main() {

}
