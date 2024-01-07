// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("START")
	t := time.Now()

	MySleep(2 * time.Second)

	fmt.Println("END", time.Since(t))
}

func MySleep(dur time.Duration) {
	ch := time.After(dur)
	<-ch
}