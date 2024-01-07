// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(Remove(arr, 2))
}

func Remove(arr []int, indx int) []int {
	return append(arr[:indx], arr[indx+1:]...)
}