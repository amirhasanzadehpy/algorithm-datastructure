package topics

import "fmt"

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
	fmt.Println(h.values)
}
