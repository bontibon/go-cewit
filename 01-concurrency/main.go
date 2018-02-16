package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func ProcessString(wg *sync.WaitGroup, s string) {
	time.Sleep(time.Millisecond * 333) // Perform some blocking operation
	fmt.Println(strings.ToUpper(s))
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	items := []string{"hotel", "alpha", "charlie", "kilo"}
	for _, item := range items {
		wg.Add(1)
		go ProcessString(&wg, item)
	}
	wg.Wait()
}
