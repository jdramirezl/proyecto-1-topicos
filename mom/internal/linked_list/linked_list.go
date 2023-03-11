package linked_list

import "errors"

const ErrEmptyList = errors.New("linked list is empty")

type LinkedList struct {
	Head   *node
	Length int
}

type node struct {
	Val  string
	Next *node
}

func (q *LinkedList) Add(val string) {
	q.Length++

	if q.Head == nil {
		q.Head = &node{Val: val}
		return
	}
	n := q.Head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &node{Val: val}
}

func (q *LinkedList) Pop() error {
	if q.Head == nil {
		return ErrEmptyList
	}
	q.Head = q.Head.Next
	q.Length--
	return nil
}

func (q *LinkedList) Peek() (string, error) {
	if q.Head == nil {
		return "", ErrEmptyList
	}
	return q.Head.Val, nil
}

func (q *LinkedList) IsEmpty() bool {
	return q.Head == nil
}

func (q *LinkedList) Size() int {
	return q.Length
}

func (q *LinkedList) Clear() {
	q.Head = nil
	q.Length = 0
}

func (q *LinkedList) Get(val string) string {
	n := q.Head
	for n != nil {
		if n.Val == val {
			return n.Val
		}
		n = n.Next
	}
	return ""
}
