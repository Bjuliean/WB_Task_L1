// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Counter struct {
	Value int
	mu    sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		Value: 0,
		mu: sync.RWMutex{},
	}
}

func (c *Counter) Inc() {
	c.mu.Lock() // для избежания race condition используем мьютекс
	defer c.mu.Unlock()
	c.Value++
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	totalInc, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}

	ct := NewCounter()

	for i := 0; i < totalInc; i++ { // запускаем заданное количество горутин
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			ct.Inc()
		}()
	}

	wg.Wait()

	fmt.Println(ct.Value)
}
