package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // 1 ~ 3
	fmt.Println(s)

	s = s[:2] // 1 ~ 1
	fmt.Println(s)

	s = s[1:] // 1
	fmt.Println(s)
}