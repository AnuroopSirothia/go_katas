package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var wg sync.WaitGroup
	inchan1 := make(chan string)
	inchan2 := make(chan string)
	quit := make(chan bool)
	wg.Add(1)

	go listen(inchan1, inchan2, quit, &wg)

	go fmt.Println("Specify input in the form of chan1=<some string>...")
	for scanner.Scan() {
		input := scanner.Text()

		switch {
		case strings.HasPrefix(input, "c1"):
			inchan1 <- input

		case strings.HasPrefix(input, "c2"):
			inchan2 <- input

		case strings.HasPrefix(input, "quit"):
			quit <- false
			wg.Wait()
			return
		}

		fmt.Println("Waiting for input...")
	}
}

func listen(inchan1 chan string, inchan2 chan string, quit chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case a := <-inchan1:
			fmt.Println("Channel 1 = ", a)
		case b := <-inchan2:
			fmt.Println("Channel 2 = ", b)
		case <-quit:
			fmt.Println("Stopping goroutine.")
			wg.Done()
			return
		}
	}
}
