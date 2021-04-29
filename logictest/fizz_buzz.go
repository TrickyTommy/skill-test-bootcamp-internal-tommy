package main

import "fmt"

func main() {
	var input [5]string
	hasil := make([]string, len(input))

	for i := 1; i < len(hasil); i++ {
		if i%3 == 0 {
			hasil = append(hasil, "fizz")
			fmt.Print(i, hasil)
		}

		if i%5 == 0 {
			hasil = append(hasil, "buzz")
			fmt.Print(i, hasil)
		}
	}
}
