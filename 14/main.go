// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func main() {
	var a int
	var b bool

	DefineType(a)
	DefineType(b)
}

func DefineType(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	case chan any:
		fmt.Println("chan")
	default:
		fmt.Println("unknown")
	}
}