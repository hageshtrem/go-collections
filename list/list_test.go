package list

import (
	"fmt"
	"testing"
)

var implementationError error = fmt.Errorf("Incorrect implementation")

func TestList(t *testing.T) {
	l := New()
	v1 := l.PushFront(1)
	v2 := l.PushFront(2)
	if l.Front() != v2 {
		t.Errorf("%v, list: [%s]", implementationError, l)
	}
	v3 := l.PushBack(3)
	if l.Back() != v3 {
		t.Errorf("%v, list: [%s]", implementationError, l)
	}
	l.Remove(v2)
	if v := l.Front(); v == nil || v != v1 {
		t.Errorf("%v, list: [%s]", implementationError, l)
	}
}
