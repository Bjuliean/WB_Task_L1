// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// проблема этого кода в том, что указатель в justString, глобальной строке,
// указывает на большую локальную переменную. из-за того, что даже после выхода
// из функции, justString продолжит ссылаться на эту переменную, сборщик мусора
// ее не очистит.

package main

import (
	"fmt"
	"strings"
)

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = strings.Clone(v[:100]) // проблему можно решить, если использовать полноценное копирование

  fmt.Println(justString)
}

func createHugeString(sz int64) string {
	var res string

	for i := int64(0); i < sz; i++ {
		res += "a"
	}
	
	return res
}

func main() {
  someFunc()
}