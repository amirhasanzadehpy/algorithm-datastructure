package topics

import (
	"fmt"
)

type heaps struct {
	values []int
}

func NewHeaps() *heaps {
	return &heaps{values: []int{}}
}

func (h *heaps) bubbleUp() {
	idx := len(h.values) - 1
	element := h.values[idx]
	for {
		parentIdx := (idx - 1) / 2
		parent := h.values[parentIdx]
		if parent >= element {
			break
		}
		h.values[parentIdx], h.values[idx] = element, parent
		idx = parentIdx
	}
}

func (h *heaps) Insert(value int) {
	h.values = append(h.values, value)
	h.bubbleUp()
}
func (h *heaps) ExtractMax() int {
	max := h.values[0]
	end := h.values[len(h.values)-1]
	// pop
	h.values = h.values[:len(h.values)-1]
	if len(h.values) > 0 {
		h.values[0] = end
		h.sinkDown()
	}

	return max
}

func (h *heaps) Print() {
	fmt.Println(h.values)
}

func (h *heaps) sinkDown() {
	idx := 0
	element := h.values[0]
	length := len(h.values)
	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2
		var leftChild, rightChild int
		var swap int

		if leftChildIdx < length {
			leftChild = h.values[leftChildIdx]
			if element < leftChild {
				swap = leftChildIdx
			}
		}

		if rightChildIdx < length {
			rightChild = h.values[rightChildIdx]
			if swap == 0 && element < rightChild || swap != 0 && rightChild > leftChild {
				swap = rightChildIdx
			}
		}

		if swap == 0 {
			break
		}

		h.values[idx], h.values[swap] = h.values[swap], h.values[idx]
		idx = swap
	}
}
