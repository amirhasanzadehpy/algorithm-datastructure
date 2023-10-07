package topics

import "fmt"

type prioritynode struct {
	value    any
	priority int
}

type priorityQueue struct {
	values []*prioritynode
}

func NewPriorityQueue() *priorityQueue {
	return &priorityQueue{}
}

func (p *priorityQueue) Print() {
	for _, value := range p.values {
		fmt.Println(value)
	}
}

func (p *priorityQueue) Enqueue(val any, priority int) {
	myNode := &prioritynode{
		priority: priority,
		value:    val,
	}

	p.values = append(p.values, myNode)
	p.bubbleUp()

}

func (p *priorityQueue) bubbleUp() {
	idx := len(p.values) - 1
	element := p.values[idx]
	for {
		parentIdx := (idx - 1) / 2
		parent := p.values[parentIdx]
		if parent.priority <= element.priority {
			break
		}
		p.values[parentIdx], p.values[idx] = element, parent
		idx = parentIdx
	}
}

func (p *priorityQueue) Dequeue() *prioritynode {
	min := p.values[0]
	end := p.values[len(p.values)-1]
	// pop
	p.values = p.values[:len(p.values)-1]
	if len(p.values) > 0 {
		p.values[0] = end
		p.sinkDown()
	}

	return min
}

func (p *priorityQueue) sinkDown() {
	idx := 0
	element := p.values[0]
	length := len(p.values)
	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2
		var leftChild, rightChild *prioritynode
		var swap int

		if leftChildIdx < length {
			leftChild = p.values[leftChildIdx]
			if element.priority > leftChild.priority {
				swap = leftChildIdx
			}
		}

		if rightChildIdx < length {
			rightChild = p.values[rightChildIdx]
			if swap == 0 && element.priority > rightChild.priority || swap != 0 && rightChild.priority < leftChild.priority {
				swap = rightChildIdx
			}
		}

		if swap == 0 {
			break
		}

		p.values[idx], p.values[swap] = p.values[swap], p.values[idx]
		idx = swap
	}
}
