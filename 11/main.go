// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

type Set struct {
	m map[int]struct{}
}

func NewSet() *Set { // конструктор
	s := Set{
		m: make(map[int]struct{}),
	}
	return &s
}

func (s *Set) Add(val int) {
	if s.Contains(val) {
		return
	}
	s.m[val] = struct{}{}
}

func (s *Set) Contains(val int) bool {
	if _, ok := s.m[val]; ok {
		return true
	}
	return false
}

func (s *Set) Intersection(other *Set) *Set { // пересечение множеств
	res := NewSet()
	for i, _ := range s.m {
		if other.Contains(i) {
			res.Add(i)
		}
	}
	return res
}

func main() {
	set1 := NewSet()
	set2 := NewSet()

	for _, v := range [...]int{1, 2, 3, 4, 5, 6} {
		set1.Add(v)
	}

	for _, v := range [...]int{10, 2, 12, 33, 4, 5, 9} {
		set2.Add(v)
	}

	intr := set1.Intersection(set2)

	fmt.Println(intr)
}