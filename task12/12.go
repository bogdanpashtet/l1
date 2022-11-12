// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

type void struct{}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	var set = make(map[string]void)

	for _, v := range sequence {
		if _, ok := set[v]; !ok {
			set[v] = void{}
		}
	}

	fmt.Printf("Many of strings: %v", set)
}
