package main

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix)-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[len(matrix)-j-1][len(matrix)-i-1]
			matrix[len(matrix)-j-1][len(matrix)-i-1] = temp
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			temp := matrix[j][i]
			matrix[j][i] = matrix[len(matrix)-j-1][i]
			matrix[len(matrix)-j-1][i] = temp
		}
	}	
}

func main() {

}
