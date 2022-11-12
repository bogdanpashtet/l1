// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

type void struct{}

func main() {
	var set1 = map[int]void{3: {}, 4: {}, 5: {}, 6: {}, 7: {}}
	var set2 = map[int]void{1: {}, 2: {}, 3: {}, 4: {}, 5: {}}
	var intersection = make(map[int]void)

	for k1 := range set1 {
		if _, ok := set2[k1]; ok {
			intersection[k1] = void{}
		}
	}

	fmt.Printf("First set: %v\nSecond set: %v\nIntersection set: %v", set1, set2, intersection)

}
