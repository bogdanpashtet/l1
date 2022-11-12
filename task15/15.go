// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }
//
// func main() {
//   someFunc()
// }

package main

import (
	"fmt"
	"strconv"
)

var justString string

func createHugeString(val int) string {
	var str string
	for i := 0; i < val; i++ {
		str += strconv.Itoa(i%10) + "."
	}
	return str
}

// так как строка - это массив байтов, то и срез будет идти не посимвольно, а побайтно
// поэтому перед тем, как брать срез, преобразуем строку в срез рун
func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
	fmt.Printf("%v", justString)
}
