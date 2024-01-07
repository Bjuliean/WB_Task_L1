// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	txt := sc.Text()

	fmt.Println(ReverseWords(txt))

}

func ReverseWords(str string) string {
	strlist := strings.Split(str, " ")

	for i, j := 0, len(strlist) - 1; i < j; i, j = i + 1, j - 1 {
		strlist[i], strlist[j] = strlist[j], strlist[i]
	}

	return strings.Join(strlist, " ")
}