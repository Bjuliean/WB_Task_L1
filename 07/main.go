// Реализовать конкурентную запись данных в map.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type db struct {
	M  map[string]int
	Mu sync.RWMutex // для избежания race condition, используем mutex
}

func (d *db) Add(i int) {
	d.Mu.Lock()
	defer d.Mu.Unlock()
	d.M[fmt.Sprintf("KeY%d", i)] = i 
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	totalRecords, err := strconv.Atoi(sc.Text()) // предлагаем ввести количество записей
	if err != nil {
		log.Fatal(err)
	}

	st := db{
		M: make(map[string]int),
		Mu: sync.RWMutex{},
	}

	wg := sync.WaitGroup{}

	for i := 0; i < totalRecords; i++ {
		wg.Add(1)
		go func (z int)  {
			defer wg.Done()
			st.Add(z) // конкурентная запись в map
		}(i)
	}

	wg.Wait() // дожидаемся горутины

	for _, i := range st.M {
		fmt.Println(i)
	}
}
