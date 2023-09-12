package main

import "fmt"

func split(sum int) (엑스, 와이 int) {
	엑스 = sum * 4 / 9
	와이 = sum - 엑스
	return
}

func main() {
	fmt.Println(split(17))
}
