package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func main() {
	nodes := make([]Node, 3)
	nodes[2] = Node{"third", nil}
	nodes[1] = Node{"second", &nodes[2]}
	nodes[0] = Node{"first", &nodes[1]}
	fmt.Println(nodes)

	for i := 0; i <= 2; i++ {
		if i == 0 {
			nodes[i].next = nil
		} else {
			nodes[i].next = &nodes[i-1]
		}
	}
	fmt.Println(nodes)
}
