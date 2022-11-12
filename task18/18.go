// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type counter struct {
	value int
	mx    sync.Mutex
}

func (T *counter) Inc(wg *sync.WaitGroup) {
	T.mx.Lock()
	T.value++
	T.mx.Unlock()
	wg.Done()
}

func main() {
	var cnt counter
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go cnt.Inc(&wg)
	}
	wg.Wait()

	fmt.Printf("Value of structure-counter: %v", cnt.value)
}
