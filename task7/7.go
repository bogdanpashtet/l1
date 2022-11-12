// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

// main1 использование Map и sync.Mutex для конкурентной записи
func main1() {
	var m = make(map[int]string)
	var mux sync.Mutex
	var wg sync.WaitGroup
	countOfValues := 15

	for i := 0; i < countOfValues; i++ {
		wg.Add(1)
		go func(key int) {
			mux.Lock()
			m[key] = fmt.Sprintf("Value № %v", key)
			mux.Unlock()
			wg.Done()
			fmt.Printf("%v\t%v\n", key, m[key])
		}(i)
	}
	wg.Wait()
}

// main2 использование sync.Map для конкурентной записи
func main2() {
	var m sync.Map
	var wg sync.WaitGroup
	countOfValues := 15

	for i := 0; i < countOfValues; i++ {
		wg.Add(1)
		go func(key int) {
			m.Store(key, fmt.Sprintf("Value № %v", key))
			wg.Done()
			v, _ := m.Load(key)
			fmt.Printf("%v\t%v\n", key, v)
		}(i)
	}
	wg.Wait()
}

func main() {
	fmt.Println("\nMap + sync.Mutex:")
	main1()
	fmt.Println("\nsync.Map:")
	main2()
}
