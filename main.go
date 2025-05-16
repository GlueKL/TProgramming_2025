package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Грачев Антон Михайлович")
	fmt.Println("Вариант 9")

	fmt.Println("y1: ", solve(1.21, 1.1, 0.09))
	fmt.Println("y2: ", solve(1.76, 1.1, 0.09))
	fmt.Println("y3: ", solve(2.53, 1.1, 0.09))
	fmt.Println("y4: ", solve(3.48, 1.1, 0.09))
	fmt.Println("y5: ", solve(4.52, 1.1, 0.09))

}

func solve(x, a, b float64) float64 {
	part1 := math.Log10(math.Pow(x, 2) - 1)
	part2 := math.Log(a*math.Pow(x, 2)-b) / math.Log(5)
	y := part1 / part2
	return y
}
