package generics

//--- StringSet

type StringSet struct {
	items map[string]int
}

func NewSet() StringSet {
	return StringSet{items: make(map[string]int)}
}

func (s *StringSet) Add(item string) {
	s.items[item] = 1
}

func (s *StringSet) Remove(item string) {
	delete(s.items, item)
}

func (s StringSet) Has(item string) (ok bool) {
	_, ok = s.items[item]
	return
}

func (s *StringSet) List() (list []string) {
	for k := range s.items {
		list = append(list, k)
	}
	return
}

//--- TSet

type TSet[K comparable, V int] struct {
	items map[K]V
}

func NewTSet[K comparable, V int]() TSet[K, V] {
	return TSet[K, V]{items: make(map[K]V)}
}

func (s *TSet[K, V]) Add(item K) {
	s.items[item] = 1
}

func (s *TSet[K, V]) Remove(item K) {
	delete(s.items, item)
}

func (s TSet[K, V]) Has(item K) (ok bool) {
	_, ok = s.items[item]
	return
}

func (s *TSet[K, V]) List() (list []K) {
	for k := range s.items {
		list = append(list, k)
	}
	return
}
