package main

import (
	"fmt"
)

// transposeMatrix выполняет транспонирование матрицы.
func transposeMatrix(matrix [][]int) [][]int {
	// Определяем количество строк и столбцов в исходной матрице.
	rows := len(matrix)
	cols := len(matrix[0])

	// Создаем новую матрицу для хранения транспонированной версии.
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	// Транспонируем матрицу, меняя местами строки и столбцы.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func main() {
	// Пример входной матрицы.
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Исходная матрица:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Транспонирование матрицы.
	transposedMatrix := transposeMatrix(matrix)

	fmt.Println("\nТранспонированная матрица:")
	for _, row := range transposedMatrix {
		fmt.Println(row)
	}
}
