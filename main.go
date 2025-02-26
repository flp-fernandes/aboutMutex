package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex

	contador := 0

	for range 10 {
		go func() {
			mutex.Lock()

			contador++

			fmt.Println("Contador:", contador)

			mutex.Unlock()
		}()
	}

	time.Sleep(time.Second)
}
