package main

import (
	"fmt"
	"strings"
)

var result string

func main() {
	str := "Ibu Ratna antar ubi"
	for _, v := range str {
		result = string(v) + result
	}
	strRev := result

	fmt.Println(strRev)

	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}
	str2 := string(rns)
	//convert slice to string///
	// str2 := strings.Join(hasil, " ")
	//check palindrome
	if strings.ToLower(str) == strings.ToLower(str2) {
		fmt.Println(str, "is Palindrome")
	} else {
		fmt.Println(str, "Not Palindrome")
	}
}
