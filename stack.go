package generics

import "golang.org/x/exp/constraints"

// -- Stack using string as element

type StringStack struct {
	items []string
}

func (s *StringStack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *StringStack) Pop() (item string) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
	}
	return
}

func (s *StringStack) Peek() (item string) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
	}
	return
}

func (s *StringStack) IsEmpty() bool {
	return len(s.items) == 0
}

// -- Stack using interface{} as element

type IStack struct {
	items []interface{}
}

func (s *IStack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *IStack) Pop() (item interface{}) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
	}
	return
}

func (s *IStack) Peek() (item interface{}) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
	}
	return
}

func (s *IStack) IsEmpty() bool {
	return len(s.items) == 0
}

// -- Stack using any as element

type Stack struct {
	items []any
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (item any) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
	}
	return
}

func (s *Stack) Peek() (item any) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
	}
	return
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// -- Stack with type parameters

type TStack[T constraints.Ordered] struct {
	items []T
}

func (s *TStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *TStack[T]) Pop() (item T) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
	}
	return
}

func (s *TStack[T]) Peek() (item T) {
	if !s.IsEmpty() {
		item = s.items[len(s.items)-1]
	}
	return
}

func (s *TStack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
