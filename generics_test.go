package generics

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	var a, b int = 1, 2
	var c, d float64 = 3.2, 4.3

	ans1 := Add(a, b)
	if ans1 != 3 && reflect.Kind(ans1) != reflect.Int {
		t.Error(ans1)
	}

	ans2 := Add(c, d)
	if ans2 != 7.5 && reflect.Kind(ans2) != reflect.Float64 {
		t.Error(ans2)
	}

	ans3 := AddNumbers(c, d)
	if ans3 != 7.5 && reflect.Kind(ans3) != reflect.Float64 {
		t.Error(ans3)
	}

}

func TestStringStack(t *testing.T) {
	stack := &StringStack{}
	stack.Push("the")
	stack.Push("quick")
	stack.Push("brown")
	stack.Push("fox")

	if stack.Peek() != "fox" {
		t.Error(stack.Peek())
	}
	if str := stack.Pop(); str != "fox" {
		t.Error(str)
	}
	if str := stack.Pop(); str != "brown" {
		t.Error(str)
	}
}

func TestAnyStack(t *testing.T) {
	stack := &Stack{}
	stack.Push("the")
	stack.Push("quick")
	stack.Push("brown")
	stack.Push(100)

	if stack.Peek() != 100 {
		t.Error(stack.Peek())
	}
	if i := stack.Pop(); i != 100 {
		t.Error(i)
	}
	if str := stack.Pop(); str != "brown" {
		t.Error(str)
	}
}

func TestTStack(t *testing.T) {
	stack := &TStack[string]{}
	stack.Push("the")
	stack.Push("quick")
	stack.Push("brown")
	stack.Push("fox")

	if stack.Peek() != "fox" {
		t.Error(stack.Peek())
	}
	if str := stack.Pop(); str != "fox" {
		t.Error(str)
	}
	if str := stack.Pop(); str != "brown" {
		t.Error(str)
	}
}

func TestTStackInt(t *testing.T) {
	stack := &TStack[int]{}
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)
	stack.Push(400)

	if stack.Peek() != 400 {
		t.Error(stack.Peek())
	}
	if i := stack.Pop(); i != 400 {
		t.Error(i)
	}
	if i := stack.Pop(); i != 300 {
		t.Error(i)
	}
}

func TestSet(t *testing.T) {
	set := NewSet()
	set.Add("the")
	set.Add("quick")
	set.Add("brown")
	set.Add("fox")
	set.Add("the")

	if !set.Has("quick") {
		t.Error("set doesn't have quick")
	}
	list := set.List()
	sort.Slice(list, func(i, j int) bool {
		return strings.Compare(list[i], list[j]) == -1
	})
	if !reflect.DeepEqual(list, []string{"brown", "fox", "quick", "the"}) {
		t.Error(list)
	}
}

func TestTSet(t *testing.T) {
	set := NewTSet[string]()
	set.Add("the")
	set.Add("quick")
	set.Add("brown")
	set.Add("fox")
	set.Add("the")

	if !set.Has("quick") {
		t.Error("set doesn't have quick")
	}
	list := set.List()
	sort.Slice(list, func(i, j int) bool {
		return strings.Compare(list[i], list[j]) == -1
	})
	if !reflect.DeepEqual(list, []string{"brown", "fox", "quick", "the"}) {
		t.Error(list)
	}
}

func TestTSetInt(t *testing.T) {
	set := NewTSet[int]()
	set.Add(100)
	set.Add(200)
	set.Add(300)
	set.Add(400)

	if !set.Has(400) {
		t.Error("set doesn't have 400")
	}
	list := set.List()
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	if !reflect.DeepEqual(list, []int{100, 200, 300, 400}) {
		t.Error(list)
	}
}
