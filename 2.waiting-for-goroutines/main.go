package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printNumber(&wg)
	go printLetter(&wg)

	wg.Wait()
}

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		// time.Sleep(100 * time.Millisecond)
	}
}

func printLetter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'j'; i++ {
		fmt.Println(string(i))
		// time.Sleep(100 * time.Millisecond)
	}
}
