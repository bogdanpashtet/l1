// Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import (
	"fmt"
	"strconv"
)

func groupValue(val float32) string {
	if -10 < val && val < 0 {
		return "-0"
	} else {
		return strconv.Itoa(int(val/10) * 10)
	}
}

func main() {
	var m = make(map[string][]float32)
	sl := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 4.9, 1.5, -2.6, -10}

	for _, v := range sl {
		group := groupValue(v)
		if val, ok := m[group]; ok {
			val = append(val, v)
			m[group] = val
		} else {
			var buffSlice = []float32{v}
			m[group] = buffSlice
		}
	}

	fmt.Printf("%v", m)
}
