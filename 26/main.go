// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func main() {
	strlist := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, v := range strlist {
		fmt.Println(IsUnique(v))
	}
}

func IsUnique(str string) bool {
	workstr := []rune(strings.ToLower(str))
	m := make(map[rune]int, len(workstr)) // для определения уникальности используем map

	for i := 0; i < len(workstr); i++ {
		m[workstr[i]]++
		
		if m[workstr[i]] > 1 { // в данном случае значение элемента map, показывает сколько раз такой ключ уже встречался
			return false
		}
	}

	return true
}