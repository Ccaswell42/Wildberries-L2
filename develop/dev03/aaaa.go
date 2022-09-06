package main

import (
	"fmt"
)

func Bublesort(num []int) []int {
	for i := 0; i < len(num)-1; i++ {
		for j := len(num) - 1; j > i; j-- {
			if num[j-1] > num[j] {
				num[j-1], num[j] = num[j], num[j-1]
			}
		}
	}
	return num
}
func main() {
	num := []int{190, 22, -9, 0, 17, 1388}
	fmt.Println(Bublesort(num))
}
