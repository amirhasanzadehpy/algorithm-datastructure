package topics

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
}

func NewBst() *bst {
	return &bst{}
}

func (b *bst) Insert(val int) {
	myNode := &node{value: val}
	if b.root == nil {
		b.root = myNode
	} else {
		current := b.root
		for {
			if val == current.value {
				return
			}
			if val < current.value {
				if current.left == nil {
					current.left = myNode
					return
				} else {
					current = current.left
				}
			} else if val > current.value {
				if current.right == nil {
					current.right = myNode
					return
				} else {
					current = current.right
				}
			}
		}
	}
}

func (b *bst) Find(value int) bool {
	if b.root == nil {
		return false
	}

	current := b.root
	found := false
	for current != nil && !found {
		if value < current.value {
			current = current.left
		} else if value > current.value {
			current = current.right
		} else {
			found = true
		}
	}

	return !found
}
