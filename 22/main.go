// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter first num")
	sc.Scan()
	a := big.Int{}
	a.UnmarshalText(sc.Bytes())

	fmt.Println("Enter second num")
	sc.Scan()
	b := big.Int{}
	b.UnmarshalText(sc.Bytes())

	fmt.Println("Enter action")
	sc.Scan()
	fmt.Println(Action(sc.Text(), a, b))

}

func Action(act string, a, b big.Int) *big.Int {
	var res big.Int

	switch act {
	case "+":
		return res.Add(&a, &b)
	case "-":
		return res.Sub(&a, &b)
	case "*":
		return res.Mul(&a, &b)
	case "/":
		return res.Div(&a, &b)
	}

	return &res
}
