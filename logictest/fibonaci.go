package main

import "fmt"

func main() {
	var input FibonaInput
	input.nFibona = []int{15, 1, 3}

	nFibonacci := input.nFibonacci()

	fmt.Println(nFibonacci)
}

type FibonaInput struct {
	nFibona []int
}

func (nf FibonaInput) nFibonacci() int {
	n := 0
	t1 := 0
	t2 := 1
	nextTerm := 0

	for _, v := range nf.nFibona {
		n += v
	}

	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
	}

	return nextTerm
}
