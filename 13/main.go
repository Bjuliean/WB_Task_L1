// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a, b := 5, 10

	fmt.Println(a, b)

	a, b = b, a

	fmt.Println(a, b)
}