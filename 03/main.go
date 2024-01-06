// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	var res int
	ch := make(chan int, len(arr))
	wg := sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ch <- (val * val)
		}(v)
	}

	wg.Wait()
	close(ch) // ожидаем завершения горутин, закрываем канал

	for v := range ch { // поскольку размеры канала и массивов совпадают, а
		res += v        // канал закрыт, range переберет все лежащие в канале значения, а затем выйдет из цикла
	}					// получив правильный ответ

	fmt.Println(res)
}