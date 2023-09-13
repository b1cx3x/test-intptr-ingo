package main

import (
	"fmt"
	"math"
)

func split(썸 int) (엑스, 와이 int) {
	엑스 = 썸 * 4 / 9
	와이 = 썸 - 엑스
	return
}

func swap(x, y string) (string, string) {
	return x + y, y + x
}

func pow(x, n, smaller float64) float64 {
	if v := math.Pow(x, n); v < smaller {
		return v
	}
	return smaller
}

func main() {
	c := "no"
	fmt.Println(split(17))
	fmt.Println(swap("하세요", "안녕"))
	fmt.Printf("The Type of c is %T\n", c)
	fmt.Println(pow(5, 2, 24.2))
}
