// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"math"
)

func main() {
	var t int64 = 0
	t = SetBit(t, 10, true)
	t = SetBit(t, 5, true)
	t = SetBit(t, 5, true)
	fmt.Printf("%b\n", t)

	fmt.Printf("%b\n", SetBit(t, 10, false))
	fmt.Printf("%b\n", SetBit(t, 10, false))
}

func SetBit(num int64, bit uint, set bool) int64 { // принимаем число, номер бита начиная с 0, и хотим ли мы установить единицу
	var i int64 = 1
	i <<= bit // побитовый сдвиг влево

	if !set {
		i ^= math.MaxInt64 // делаем реверс i, теперь на месте нулей будут единицы, и наоборот
		num &= i // побитовое и, ноль убирает бит с заданной позиции
	} else {
		num |= i // побитовое или, на заданной позиции будет единица
	}

	return num
}