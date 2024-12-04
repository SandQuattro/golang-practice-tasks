package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// база рекурсии, если вершина nil, вернем 0
	if root == nil {
		return 0
	}

	// считаем максимум пути для левой части дерева и для правой
	// делая рекурсивный обход
	maxLeftPath := maxPathSum(root.left)
	maxRightPath := maxPathSum(root.right)

	// итоговый результат будет суммы всех путей + значение самой вершины root
	return max(maxLeftPath, maxRightPath) + root.val
}

func main() {
	// Тест 1: Простое дерево
	tree1 := &TreeNode{val: 1,
		left:  &TreeNode{val: 3},
		right: &TreeNode{val: 2}}
	fmt.Println("Тест 1:", maxPathSum(tree1)) // Ожидаемый результат: 4

	// Тест 2: Левый путь больше
	tree2 := &TreeNode{val: 1,
		left:  &TreeNode{val: 5},
		right: &TreeNode{val: 2}}
	fmt.Println("Тест 2:", maxPathSum(tree2)) // Ожидаемый результат: 6

	// Тест 3: Правый путь больше
	tree3 := &TreeNode{val: 1,
		left:  &TreeNode{val: 2},
		right: &TreeNode{val: 5}}
	fmt.Println("Тест 3:", maxPathSum(tree3)) // Ожидаемый результат: 6

	// Тест 4: Отрицательные значения
	tree4 := &TreeNode{val: 1,
		left:  &TreeNode{val: -3},
		right: &TreeNode{val: 2}}
	fmt.Println("Тест 4:", maxPathSum(tree4)) // Ожидаемый результат: 3

	// Тест 5: Глубокое дерево
	tree5 := &TreeNode{val: 1,
		left: &TreeNode{val: 2,
			left: &TreeNode{val: 4}},
		right: &TreeNode{val: 3}}
	fmt.Println("Тест 5:", maxPathSum(tree5)) // Ожидаемый результат: 7

	// Тест 6: Единственный узел
	tree6 := &TreeNode{val: 5}
	fmt.Println("Тест 6:", maxPathSum(tree6)) // Ожидаемый результат: 5

	// Тест 7: Пустое дерево
	var tree7 *TreeNode
	fmt.Println("Тест 7:", maxPathSum(tree7)) // Ожидаемый результат: 0
}
