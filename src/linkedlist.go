package datastructs

import "fmt"

type LinkedList[T any] struct {
    head *node[T]
    tail *node[T]
    length int
}

func createnode[T any](value T) *node[T] {
    return &node[T]{value, nil};
}

func (ll *LinkedList[T]) Append(value T) (head *node[T]) {
    node := createnode(value);
    ll.length++
    if (ll.head == nil && ll.tail == nil) {
        ll.head = node
        ll.tail = node
        return ll.head
    }
    ll.tail.next = node;
    ll.tail = node;
    return ll.head;
}

func (ll *LinkedList[T]) InsertAt(value T, index int) {}
func (ll *LinkedList[T]) DeleteAt(value T, index int) {}
func (ll *LinkedList[T]) Prepend(value T, index int) {}
func (ll *LinkedList[T]) At(index int) {} 
func (ll *LinkedList[T]) Pop() {} 

func (ll *LinkedList[T]) PrintList(){
    if (ll.length == 0) {
        return;
    }
    fmt.Printf("LL H( %v ), T( %v ), L( %d ) => ", ll.head.value, ll.tail.value, ll.length)
    tmp := ll.head
    for range ll.length {
        if tmp.next ==  nil {
            fmt.Printf("( %v )\n", tmp.value)
            return;
        }
        fmt.Printf("( %v ) -> ", tmp.value)
        tmp = tmp.next
    }
}
