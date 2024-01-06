// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	size, err := strconv.Atoi(sc.Text()) // вводим желаемый размер массива
	if err != nil {
		log.Fatal(err)
	}

	arr := make([]int, size)
	fmt.Printf("GENERATED ARRAY: ")
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size) // заполняем массив случайными числами
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")

	in := make(chan int)
	out := make(chan int)
	wg := sync.WaitGroup{} // создаем каналы и waitgroup

	go Count(in, out)
	wg.Add(1)
	go Write(out, &wg) // в waitgroup добавляем только Write, поскольку 
					   // данные идут конвейером, а Write завершающая функция

	go func ()  {
		defer close(in)

		for _, i := range arr {
			in <- i
		}
	}()

	wg.Wait()
}

func Count(in chan int, out chan int) {
	for v := range in {
		out <- (v * 2)
	}
	close(out)
}

func Write(out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("OUTPUT ARRAY: ")
	for v := range out {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}