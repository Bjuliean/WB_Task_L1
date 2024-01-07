// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	arrLen, err := strconv.Atoi(sc.Text()) // ввод длины массива
	if err != nil {
		log.Fatal(err)
	}

	arr := []int{}

	for i := 0; i < arrLen; i++ {
		arr = append(arr, rand.Intn(10000))
	}

	fmt.Println("ORIGINAL", arr)

	st := time.Now()

	qsort(arr)

	fmt.Println("SORTED", arr)
	
	fmt.Println("FINISHED AFTER ", time.Since(st))
}

func qsort(arr []int) {
	s, e := 0, len(arr) - 1

	sortProccess(arr, s, e)
}

func sortProccess(arr []int, s, e int) {
	if s < e {
		l, r := s, e
		p := arr[((l + r) / 2)] // опорная точка будет посередине

		for l < r {
			for arr[l] < p {
				l++
			}

			for arr[r] > p {
				r--
			}

			if l <= r {
				arr[l], arr[r] = arr[r], arr[l]
				l++
				r--
			}
		}
		
		sortProccess(arr, s, r) // также сортируем сформировавшиеся подмассивы
		sortProccess(arr, l, e)
	}
}
