// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	
	wg.Add(1)
	ch1 := make(chan int)
	go f(ch1, &wg) // ожидание первого входящего значения

	wg.Add(1)
	ch2 := make(chan int)
	go s(ch2, &wg) // ожидание первого входящего значения через select

	wg.Add(1)
	ch3 := make(chan int)
	go t(ch3, &wg) // выход из цикла range при закрытии канала

	wg.Add(1)
	ch4 := time.After(300 * time.Millisecond)
	go fo(ch4, &wg) // по истечении времени

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go fi(ctx, &wg) // через контекст

	ch1 <- 1
	ch2 <- 2
	close(ch3)
	cancel()

	wg.Wait()
}

func fi(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("FIFTH STOPPED")
			return
		}
	}
}

func fo(ch <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	
	<-ch

	fmt.Println("FOURTH STOPPED")
}

func t(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		_ = v
		time.Sleep(1 * time.Second)
	}
	fmt.Println("THIRD STOPPED")
}

func s(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			fmt.Println("SECOND STOPPED")
			return
		}
	}
}

func f(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		<-ch
		fmt.Println("FIRST STOPPED")
		return
	}
}