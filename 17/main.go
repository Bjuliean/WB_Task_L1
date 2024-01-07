// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8} // бинарный поиск работает с отсортированными массивами

	fmt.Println(binSearch(arr, 4))
	fmt.Println(binSearch(arr, 8))
	fmt.Println(binSearch(arr, 0))
}

func binSearch(arr []int, target int) int {
	l, r, res := 0, len(arr) - 1, 0

	if r <= 0 {
		return 0
	}

	for l <= r {
		p := (l+r)/2 // индекс серединного элемента

		if arr[p] == target {
			return p 
		}

		if arr[p] < target { // если элемент посередине меньше цели, то все, что левее - нас не интересует
			l = p + 1
		}

		if arr[p] > target { // если элемент посередине больше цели, то все, что правее - нас не интересует
			r = p - 1
		}
	}

	return res 
}