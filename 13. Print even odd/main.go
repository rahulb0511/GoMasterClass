package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(ch1 chan string, ch2 chan string) {
		for val := 1; val <= 20; val++ { //prints even
			if val%2 == 0 {
				<-ch1
				//recieve a msg
				fmt.Print(val)
				ch2 <- "print odd"
				//send a message in ch
			}
		}
		wg.Done()
	}(ch1, ch2)

	go func(ch1 chan string, ch2 chan string) { //prints odd
		for val := 1; val <= 20; val++ {
			if val%2 != 0 {
				<-ch2
				//recieve a msg
				fmt.Print(val)
				//send a msg
				ch1 <- "print even"
			}
		}
		wg.Done()
	}(ch1, ch2)
	ch2 <- "start"
	wg.Wait()
}
