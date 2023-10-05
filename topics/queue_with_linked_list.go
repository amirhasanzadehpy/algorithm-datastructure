package topics

import "fmt"

type NodeQueue struct {
	val  interface{}
	next *NodeQueue
}

type queue struct {
	first *NodeQueue
	last  *NodeQueue
	Size  int
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) Enqueue(value interface{}) {
	newNode := &NodeQueue{val: value}

	if q.first == nil {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}

	q.Size++
}

func (q *queue) Dequeue() interface{} {
	if q.Size == 0 {
		return nil
	}

	temp := q.first

	if q.first == q.last {
		q.last = nil
	}

	q.first = temp.next
	q.Size--
	return temp.val
}

func (q *queue) Print() {
	current := q.first
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}
