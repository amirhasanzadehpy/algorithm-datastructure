package topics

import "fmt"

type Item struct {
	val  interface{}
	next *Item
}

type stack struct {
	first *Item
	last  *Item
	Size  int
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Push(value interface{}) {
	newItem := &Item{val: value}
	if s.first == nil {
		s.first = newItem
		s.last = newItem
	} else {
		newItem.next = s.first
		s.first = newItem
	}

	s.Size++
}

func (s *stack) Pop() interface{} {
	if s.Size == 0 {
		return nil
	}

	temp := s.first

	if s.first == s.last {
		s.last = nil
	}
	s.first = temp.next

	s.Size--
	return temp.val
}

func (s *stack) Print() {
	current := s.first
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}
