package generics

import "testing"

func BenchmarkStackString(b *testing.B) {
	b.ResetTimer()
	stack := &StringStack{}
	for i := 0; i < b.N; i++ {
		stack.Push("the")
		stack.Push("quick")
		stack.Push("brown")
		stack.Push("fox")
		stack.Peek()
		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Pop()
	}
}

func BenchmarkStackAny(b *testing.B) {
	b.ResetTimer()
	stack := &Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push("the")
		stack.Push("quick")
		stack.Push("brown")
		stack.Push("fox")
		stack.Peek()
		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Pop()
	}
}

func BenchmarkStackInterface(b *testing.B) {
	b.ResetTimer()
	stack := &IStack{}
	for i := 0; i < b.N; i++ {
		stack.Push("the")
		stack.Push("quick")
		stack.Push("brown")
		stack.Push("fox")
		stack.Peek()
		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Pop()
	}
}

func BenchmarkStackTypeParameter(b *testing.B) {
	b.ResetTimer()
	stack := &TStack[string]{}
	for i := 0; i < b.N; i++ {
		stack.Push("the")
		stack.Push("quick")
		stack.Push("brown")
		stack.Push("fox")
		stack.Peek()
		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Pop()
	}
}

func BenchmarkSetString(b *testing.B) {
	b.ResetTimer()
	set := NewSet()
	for i := 0; i < b.N; i++ {
		set.Add("the")
		set.Add("quick")
		set.Add("brown")
		set.Add("fox")
		set.Add("the")
		set.Remove("fox")
		set.Remove("quick")
	}
}

func BenchmarkSetGeneric(b *testing.B) {
	b.ResetTimer()
	set := NewTSet[string]()
	for i := 0; i < b.N; i++ {
		set.Add("the")
		set.Add("quick")
		set.Add("brown")
		set.Add("fox")
		set.Add("the")
		set.Remove("fox")
		set.Remove("quick")
	}
}
