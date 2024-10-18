package datastructs

import (
	"errors"
	"fmt"
)

type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func createNode[T any](value T) *node[T] {
	return &node[T]{value, nil}
}

func (ll *LinkedList[T]) Append(value T) (head *node[T]) {
	node := createNode(value)
	ll.length++
	if ll.head == nil && ll.tail == nil {
		ll.head = node
		ll.tail = node
		return ll.head
	}
	ll.tail.next = node
	ll.tail = node
	return ll.head
}

func (ll *LinkedList[T]) Delete(value T, comp func(T, T) bool) bool {
	tmp := ll.head
	if comp(tmp.value, value) {
		ll.head = tmp.next
		tmp.next = nil
		return true
	}
	for range ll.length {
		if tmp.next == nil {
			break
		}
		if comp(tmp.next.value, value) {
			ll.length--
			if tmp.next == ll.tail {
				tmp.next = nil
				ll.tail = tmp
				return true
			}
			tmp = tmp.next.next
			tmp.next.next = nil
			return true

		}
		tmp = tmp.next
	}
	return false
}

func (ll *LinkedList[T]) Prepend(value T) (head *node[T]) {
	node := createNode(value)
	ll.length++
	if ll.head == nil && ll.tail == nil {
		ll.head = node
		ll.tail = node
		return ll.head
	}
	tmp := ll.head
	node.next = tmp
	ll.head = node
	return node
}

// Returns error if passed index is grater than the list length or return the value stored at [index]
func (ll *LinkedList[T]) At(data *T, index int) (err error) {
	if index > ll.length || ll.length == 0 {
		return errors.New("INDEX_OUT_OF_RANGE")
	}
	tmp := ll.head
	for i := 0; i < ll.length; i++ {
		if i == index {
			*data = tmp.value
			return nil
		}
		tmp = tmp.next
	}
	return nil
}

func (ll *LinkedList[T]) Pop(data *T) error {
	if ll.head == nil && ll.tail == nil && ll.length == 0 {
		return errors.New("LIST_EMPTY")
	}
    ll.length--;
	tmp := ll.head
	ll.head = tmp.next
	*data = tmp.value
	return nil
}

func (ll *LinkedList[T]) PrintList() {
	if ll.length == 0 {
		return
	}
	fmt.Printf("LL H( %v ), T( %v ), L( %d ) => ", ll.head.value, ll.tail.value, ll.length)
	tmp := ll.head
	for range ll.length {
		if tmp.next == nil {
			fmt.Printf("( %v )\n", tmp.value)
			return
		}
		fmt.Printf("( %v ) -> ", tmp.value)
		tmp = tmp.next
	}
}
