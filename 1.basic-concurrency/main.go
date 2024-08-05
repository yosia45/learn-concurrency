package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumber()
	go printLetter()

	time.Sleep(2 * time.Second)
}

func printNumber() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetter() {
	for i := 'a'; i <= 'j'; i++ {
		fmt.Println(string(i))
		time.Sleep(100 * time.Millisecond)
	}
}
