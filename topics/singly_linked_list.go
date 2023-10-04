package topics

import "fmt"

type Node struct {
	val  string
	next *Node
}

type linkedList struct {
	head   *Node
	tail   *Node
	Length int
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (ll *linkedList) Push(value string) {
	myNode := &Node{
		val: value,
	}
	if ll.head == nil {
		ll.head = myNode
		ll.tail = ll.head
	} else {
		ll.tail.next = myNode
		ll.tail = myNode
	}

	ll.Length++
}

func (ll *linkedList) Print() {
	current := ll.head
	for current != nil {

		fmt.Println(current.val)
		current = current.next
	}
}

func (ll *linkedList) Pop() {
	if ll.head == nil {
		return
	}

	current := ll.head
	newTail := current
	for current.next != nil {
		newTail = current
		current = current.next
	}

	newTail.next = nil
	ll.tail = newTail
	ll.Length--
	if ll.Length == 0 {
		ll.head = nil
		ll.tail = nil
	}
}

func (ll *linkedList) Shift() {
	if ll.head == nil {
		return
	}
	ll.head = ll.head.next
	ll.Length--
	if ll.Length == 0 {
		ll.tail = nil
	}
}

func (ll *linkedList) Unshift(value string) {
	newNode := &Node{val: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}

	ll.Length++
}

func (ll *linkedList) Get(index int) *Node {
	if index < 0 || index >= ll.Length {
		return nil
	}

	current := ll.head
	counter := 0
	for counter != index {
		current = current.next
		counter++
	}

	return current
}

func (ll *linkedList) Set(index int, value string) bool {
	foundNode := ll.Get(index)
	if foundNode == nil {
		return false
	}

	foundNode.val = value
	return true
}

func (ll *linkedList) Insert(index int, value string) bool {

	if index < 0 || index > ll.Length {
		return false
	}

	if index == ll.Length {
		ll.Push(value)
		return true
	}

	if index == 0 {
		ll.Unshift(value)
		return true
	}

	newNode := &Node{val: value}

	prev := ll.Get(index - 1)
	temp := prev.next
	newNode.next = temp
	prev.next = newNode

	return true

}

func (ll *linkedList) Remove(index int) bool {
	if index < 0 || index >= ll.Length {
		return false
	}

	if index == ll.Length-1 {
		ll.Pop()
		return true
	}

	if index == 0 {
		ll.Shift()
		return true
	}

	prevNode := ll.Get(index - 1)

	prevNode.next = prevNode.next.next

	return true

}

func (ll *linkedList) Reverse() {
	node := ll.head
	ll.head = ll.tail
	ll.tail = node

	var next, prev *Node
	for i := 0; i < ll.Length; i++ {
		next = node.next
		node.next = prev
		prev = node
		node = next
	}
}
