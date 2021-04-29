package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input int
	var hasil []string
	input = 5

	for i := 1; i <= input; i++ {
		if i%3 == 0 && i%5 == 0 {
			hasil = append(hasil, "Fizzbuzz")
			continue
		}

		if i%5 == 0 {
			hasil = append(hasil, "buzz")

			continue
		}
		if i%3 == 0 {
			hasil = append(hasil, "Fizz")

			continue
		}
		hasil = append(hasil, strconv.Itoa(i))
	}
	fmt.Println(hasil)
	str := strings.Join(hasil, " ")
	fmt.Println(str)

}
