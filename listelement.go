package lf

import (
	"sync/atomic"
	"unsafe"
)

type ListElement struct {
	next     unsafe.Pointer 
	value    unsafe.Pointer
}

func (e *ListElement) Value() (value interface{}) {
	return *(*interface{})(atomic.LoadPointer(&e.value))
}

// Next returns the item on the right.
func (e *ListElement) Next() *ListElement {
	return (*ListElement)(atomic.LoadPointer(&e.next))
}

func (e *ListElement) Set(value unsafe.Pointer) {
	atomic.StorePointer(&e.value, value)
}

func (e *ListElement) CAS(from interface{}, to unsafe.Pointer) bool {
	old := atomic.LoadPointer(&e.value)
	if *(*interface{})(old) != from {
		return false
	}
	return atomic.CompareAndSwapPointer(&e.value, old, to)
}
