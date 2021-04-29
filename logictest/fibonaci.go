package main

import "fmt"

func main() {
	input := []int{15, 1, 3}

	n := 0
	param1 := 0
	param2 := 1
	m := 0

	for _, v := range input {
		n += v
	}
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", param1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", param2)
			continue
		}
		m = param1 + param2
		param1 = param2
		param2 = m
	}
	fmt.Println(m)
}
