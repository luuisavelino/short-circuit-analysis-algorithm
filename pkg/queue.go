package main

import "fmt"

func main() {
	queue := Queue{}
	queue.Enqueue("01")
	queue.Enqueue("02")
	queue.Enqueue("03")
	queue.Enqueue("04")
	queue.Enqueue("05")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}

type Queue struct {
	Head *Node
}

func (q *Queue) Enqueue(name string) {
	node := Node{Val: name}

	if q.Head == nil {
		q.Head = &node
	} else {
		curr := q.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = &node
	}
}

func (q *Queue) Dequeue() string {
	if q.Head == nil {
		return ""
	}

	node := q.Head
	q.Head = q.Head.Next

	return node.Val
}

type Node struct {
	Val string
	Next *Node
}