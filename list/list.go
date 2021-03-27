package list

import (
	"fmt"
	"strings"
)

// Element is an element of a linked list.
type Element struct {
	prev, next *Element
	list       *List
	// The value stored with this element.
	Value interface{}
}

// Next returns next element of list or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns previous element of list or nil.
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// String returns a string representation of element.
func (e *Element) String() string {
	return fmt.Sprintf("%v", e.Value)
}

// List represents a doubly linked lis
type List struct {
	root Element
	len  int
}

// New returns an initialized list.
func New() *List {
	return (&List{}).init()
}

func (l *List) init() *List {
	l.root.prev = &l.root
	l.root.next = &l.root
	l.root.list = l
	l.len = 0
	return l
}

func (l *List) insertValueAt(v interface{}, at *Element) *Element {
	e := &Element{
		prev:  at,
		next:  at.next,
		list:  l,
		Value: v,
	}
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

// PushFront inserts a new element e with value v at the front of list and returns e.
func (l *List) PushFront(v interface{}) *Element {
	return l.insertValueAt(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list and returns e.
func (l *List) PushBack(v interface{}) *Element {
	return l.insertValueAt(v, l.root.prev)
}

// Back returns the last element of list or nil if the list is empty.
func (l *List) Back() *Element {
	return l.root.prev
}

// Front returns the first element of list or nil if the list is empty.
func (l *List) Front() *Element {
	return l.root.next
}

// Len returns the number of elements of list.
func (l *List) Len() int {
	return l.len
}

// Remove removes e from l if e is an element of list.
// The element must not be nil.
func (l *List) Remove(e *Element) {
	if e.list != l {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	e.Value = nil
}

// String returns a string representation of list.
func (l *List) String() string {
	s := strings.Builder{}
	for e := l.Front(); e != nil; e = e.Next() {
		s.WriteString(e.String())
	}
	return s.String()
}
