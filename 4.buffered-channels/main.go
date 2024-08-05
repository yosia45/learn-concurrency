package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	numberCh := make(chan int, 5)
	letterCh := make(chan string, 5)

	wg.Add(4)

	go produceNumber(numberCh, &wg)
	go produceLetter(letterCh, &wg)
	go consumeNumber(numberCh, &wg)
	go consumeLetter(letterCh, &wg)

	wg.Wait()
}

func produceNumber(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func produceLetter(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'j'; i++ {
		ch <- string(i)
	}
	close(ch)
}

func consumeNumber(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range ch {
		fmt.Println(number)
	}
}

func consumeLetter(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for letter := range ch {
		fmt.Println(letter)
	}
}
