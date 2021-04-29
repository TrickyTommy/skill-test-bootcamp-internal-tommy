package main

import (
	"fmt"
	"strings"
)

var result string

func main() {
	str := "I am A Great human"
	hasil := []string{}
	strArray := strings.Fields(str)
	var satuan string
	for _, f := range strArray {
		abjad := f
		for _, c := range abjad {
			result = string(c) + result
		}
		satuan = result
		result = ""
		hasil = append(hasil, satuan)
	}
	//convert slice to string///
	str2 := strings.Join(hasil, " ")
	fmt.Println(str2)
}
