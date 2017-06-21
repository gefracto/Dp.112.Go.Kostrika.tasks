package task4

import (
	"strconv"
	"fmt"
)

func reverseString(s string) string {
	var newString string = ""
	for i := len(s) - 1; i >= 0; i-- {
		newString += string(s[i])
	}
	return newString
}

func isPal(s string) (b bool) {
	b = false
	if s == reverseString(s) && len(s) > 1 {
		b = true
	}
	return b
}

func findLargest(sl []string) []string {
	var largest []string
	var largestSize int = 1
	for _, v := range sl {
		if len(v) > largestSize {
			largestSize = len(v)
		}
	}
	for _, v := range sl {
		if len(v) == largestSize {
			largest = append(largest, v)
		}
	}
	return largest
}

func DoTask4(num int) {
	original := strconv.Itoa(num)
	reversed := reverseString(original)
	var largestPalindrome int
	var ok bool
	var sliceOfPalindromes []string
	var largestPalindromes []string
	if num <= 0 {
		fmt.Println("Ошибка!\nЗадайте число больше нуля")
	} else {
		for i1, v1 := range original {
			for i2, v2 := range reversed {
				if i1+i2 < len(original) && v2 == v1 && isPal(original[i1:len(original)-i2]) {
					sliceOfPalindromes = append(sliceOfPalindromes, string(original[i1:len(original)-i2]))
					ok = true
				}
			}
		}
		if ok {
			largestPalindromes = append(largestPalindromes, findLargest(sliceOfPalindromes)...)
		}
		for _, v := range largestPalindromes {
			tmp, _ := strconv.Atoi(v)
			if largestPalindrome < tmp {
				largestPalindrome = tmp
			}
		}
		fmt.Println(largestPalindrome)
	}
}
