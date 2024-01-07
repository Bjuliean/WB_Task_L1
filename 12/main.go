// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

type USet struct {
	m map[string]struct{} // создаем сет для типа string
}

func NewSet() *USet {
	return &USet{
		m: make(map[string]struct{}),
	}
}

func (u *USet) Add(val string) {
	if u.Contains(val) {
		return
	}
	u.m[val] = struct{}{}
}

func (u *USet) Contains(val string) bool {
	if _, ok := u.m[val]; ok {
		return true
	}
	return false
}

func main() {
	nw := NewSet()

	for _, v := range [...]string{"cat", "cat", "dog", "cat", "tree"} { // заполняем сет
		nw.Add(v)
	}

	fmt.Println(nw)
}