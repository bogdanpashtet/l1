// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
)

func main() {
	var s = []interface{}{3, 5.67, "str", 4 + 5i, make(chan int), struct{}{}, make(chan int64)}

	for _, val := range s {
		fmt.Printf("Type of %v is %T\n", val, val)
		//fmt.Printf("Type of %v is %v\n", val, reflect.TypeOf(val))
	}
}
