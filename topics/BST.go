package topics

type node struct {
	value int
	left  *node
	right *node
}

type queueTree struct {
	items []*node
}

type bst struct {
	root *node
}

func newQueue() *queueTree {
	return &queueTree{}
}

func (q *queueTree) enqueue(item *node) {
	q.items = append(q.items, item)
}

func (q *queueTree) dequeue() *node {
	if q.isEmpty() {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *queueTree) isEmpty() bool {
	return len(q.items) == 0
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

func (b *bst) BFS() []int {
	var node *node
	node = b.root
	data := []int{}
	q := newQueue()
	q.enqueue(node)
	for !q.isEmpty() {
		node = q.dequeue()
		data = append(data, node.value)
		if node.left != nil {
			q.enqueue(node.left)
		}

		if node.right != nil {
			q.enqueue(node.right)
		}
	}
	return data
}

func (b *bst) DFSPreOrder() []int {
	data := []int{}
	current := b.root

	var traverse func(node *node)
	traverse = func(node *node) {
		data = append(data, node.value)
		if node.left != nil {
			traverse(node.left)
		}

		if node.right != nil {
			traverse(node.right)
		}
	}

	traverse(current)

	return data
}
