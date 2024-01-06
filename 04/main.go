// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора
// количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

type WorkerKit struct {
	ID    int
	Tasks chan int
	Wg    *sync.WaitGroup
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	totalWorkers, err := strconv.Atoi(sc.Text()) // сохраняем количество воркеров
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	ch := make(chan int) // канал для данных

	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)
		go Worker(WorkerKit{
			ID:    i,
			Tasks: ch,
			Wg:    &wg,
		})
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for {
			ch <- rand.Int()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	<-c       // слушаем сигнал ос
	close(ch) // как только сигнал пришел, закрываем канал, заставляя горутины выйти из цикла
	wg.Wait() // дожидаемся все горутины
}

func Worker(wk WorkerKit) {
	defer wk.Wg.Done()

	for v := range wk.Tasks {
		fmt.Printf("Worker %d doing %d\n", wk.ID, v)
	}
	fmt.Printf("Worker %d finished\n", wk.ID)
}
