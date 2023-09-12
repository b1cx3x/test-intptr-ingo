package main

import "fmt"

func split(sum int) (엑스, 와이 int) {
	엑스 = sum * 4 / 9
	와이 = sum - 엑스
	return
}

func swap(x, y string) (string, string) {
	return x + y, y + x
}

func pow(x, n, lim float64) float64 {
	if v < lim {
    	return v
    }
    return lim
}

func main() {
	c := "no"
	fmt.Println(split(17))
	fmt.Println(swap("하세요", "안녕"))
	fmt.Printf("The Type of c is %T\n", c)
	fmt.Printf(pow(5, 2, 25.1))
}
