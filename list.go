package lf

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type List struct {
	head  *ListElement
	tail  *ListElement
}

func NewList() *List {
	return &List{head: &ListElement{}, tail: &ListElement{}}
}

func (l *List) Add(value interface{}) {
	q := &ListElement{}
	q.Set((unsafe.Pointer)(&value))
	q.next = nil

	var p *ListElement
	var success bool
	for {
		p = l.tail
		success = atomic.CompareAndSwapPointer(&p.next, nil, unsafe.Pointer(q))

		if (!success) {
			tail := (*unsafe.Pointer)(unsafe.Pointer(&l.tail))
			atomic.CompareAndSwapPointer(tail, unsafe.Pointer(p), p.next)
		}

		if (success) {
			break
		}
	}

	tail := (*unsafe.Pointer)(unsafe.Pointer(&l.tail))
	atomic.CompareAndSwapPointer(tail, unsafe.Pointer(p), unsafe.Pointer(q))
}

func (l *List) Delete(value interface{}) {
	//
}

func (l *List) PrintTail() {
	fmt.Println(l.tail.Value())
}