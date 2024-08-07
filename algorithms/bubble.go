package main

import (
	"fmt"
)

// моя реализация пузырьковой сортировки
func bubbleSort1(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			// меняя знак с > на < мы сортируем в порядке убывания
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Исходный массив:", arr)

	bubbleSort1(arr)

	fmt.Println("Отсортированный массив:", arr)

	bubbleSort1(nil)
	bubbleSort1([]int{})
}
