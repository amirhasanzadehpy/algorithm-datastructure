package topics

type Stack struct {
	items []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(value interface{}) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.items[len(s.items)-1]
}
