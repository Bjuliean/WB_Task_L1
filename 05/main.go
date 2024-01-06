// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	seconds, err := strconv.Atoi(sc.Text()) // запрашиваем количество секунд
	if err != nil {
		log.Fatal(err)
	}
	_ = seconds
	ch := make(chan int)
	timeout := time.After(time.Duration(seconds) * time.Second) // создаем канал, в котором засекаем время
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case ch <- rand.Int():
				time.Sleep(500 * time.Millisecond)
			case <-timeout: // как только истекает время, закрываем канал, output цикл будет вынужден прерваться
				fmt.Println("TIME OUT")
				close(ch)
				return
			}
		}
	}()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}