package main

import (
	"fmt"
	"math"
)

const (
	a = 1.1
	b = 0.09
)

func main() {

	fmt.Println("Грачев Антон Михайлович")
	fmt.Println("Вариант 9")
	fmt.Println()

	fmt.Println("Задача А")
	arr := []float64{}
	for i := 1.2; i <= 2.2; i += 0.2 {
		arr = append(arr, solve(i, a, b))
	}
	fmt.Println("Массив y: ", arr)

	fmt.Println()

	fmt.Println("Задача B")
	var arrayB []float64 = []float64{1.21, 1.76, 2.53, 3.48, 4.52}
	arrB := solveArray(arrayB)
	fmt.Println("Массив y: ", arrB)

}

func solveArray(arrayX []float64) []float64 {
	arr := []float64{}
	for _, x := range arrayX {
		arr = append(arr, solve(x, a, b))
	}
	return arr
}

func solve(x, a, b float64) float64 {
	part1 := math.Log10(math.Pow(x, 2) - 1)
	part2 := math.Log(a*math.Pow(x, 2)-b) / math.Log(5)
	y := part1 / part2
	return y
}
