// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	txt := sc.Text()

	fmt.Println(Reverse(txt))
}

func Reverse(str string) string {
	res := []rune(str)

	for i, j := 0, len(res) - 1; i < j; i, j = i + 1, j - 1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}