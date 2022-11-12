// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Start of program. Sleep after 3 seconds")
	start := time.Now()
	Sleep(3 * time.Second)
	fmt.Printf("Time has passed: %v", time.Since(start))
}
