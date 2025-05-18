package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 1, 2, 5, 5, 5, 5, 9, 9, 9, 9, 9}

	// Считаем частоты
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	// Уникальные значения для сортировки
	unique := make([]int, 0, len(freq))
	for num := range freq {
		unique = append(unique, num)
	}

	// Сортировка по убыванию частоты (и по возрастанию числа при равной частоте)
	sort.Slice(unique, func(i, j int) bool {
		if freq[unique[i]] == freq[unique[j]] {
			return unique[i] < unique[j]
		}
		return freq[unique[i]] > freq[unique[j]]
	})

	// Строим отсортированный массив
	sorted := make([]int, 0, len(arr))
	for _, num := range unique {
		for i := 0; i < freq[num]; i++ {
			sorted = append(sorted, num)
		}
	}

	fmt.Println(sorted)
}
