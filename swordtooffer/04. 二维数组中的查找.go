package main

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0 {
			continue
		}
		if matrix[i][len(matrix[i])-1] < target {
			continue
		}
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
