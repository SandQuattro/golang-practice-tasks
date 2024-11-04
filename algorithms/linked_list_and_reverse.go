package main

import "fmt"

// Node представляет один узел односвязного списка
type Node struct {
	value int
	next  *Node
}

// LinkedList представляет односвязный список
type LinkedList struct {
	head *Node
}

// Add добавляет новый элемент в конец списка
func (l *LinkedList) Add(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Print выводит элементы списка
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

// Reverse обращает односвязный список
func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.head
	for current != nil {
		next := current.next // сохраняем следующий узел
		current.next = prev  // переворачиваем ссылку
		prev = current       // передвигаем prev на один узел вперед
		current = next       // передвигаем current на один узел вперед
	}
	l.head = prev // обновляем голову списка
}

func main() {
	list := LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	fmt.Println("Исходный список:")
	list.Print()

	list.Reverse()
	fmt.Println("Обращённый список:")
	list.Print()
}
