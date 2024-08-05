package main

import "fmt"

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go func() {
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
				evenCh <- i
			} else {
				oddCh <- i
			}
		}
		close(evenCh)
		close(oddCh)
	}()

	for {
		select {
		case num, ok := <-oddCh:
			if !ok {
				oddCh = nil
				continue
			}
			fmt.Printf("Received an odd number: %d\n", num)
		case num, ok := <-evenCh:
			if !ok {
				evenCh = nil
				continue
			}
			fmt.Printf("Received an even number: %d\n", num)
		}
		if oddCh == nil && evenCh == nil {
			break
		}
	}
}
