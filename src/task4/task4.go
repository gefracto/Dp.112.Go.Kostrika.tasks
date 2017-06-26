package task4

import (
	"errors"
	"strconv"
)

// make it using ints and nit strings!

type T4 struct {
	Number int
}

func (T *T4) Dotask4() (err error, data int) {
	return Dotask(T.Number)
}

func reverseString(s string) string {
	var newString string = ""
	for i := len(s) - 1; i >= 0; i-- {
		newString += string(s[i])
	}
	return newString
}

func isPal(s string) (b bool) {
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

func Dotask(num int) (err error, data int) {
	var ok bool
	ABC := strconv.Itoa(num)
	CBA := reverseString(ABC)
	var largestPalindrome int
	var palindromes []string
	var largestPalindromes []string
	if num <= 0 {
		return errors.New("\nОшибка!\nЗадайте число больше нуля"), data

	}

	for i1, v1 := range ABC {
		for i2, v2 := range CBA {
			if i1+i2 < len(ABC) && v2 == v1 && isPal(ABC[i1:len(ABC)-i2]) {
				palindromes = append(palindromes, string(ABC[i1:len(ABC)-i2]))

				ok = true
			}
		}
	}

	if ok {
		largestPalindromes = append(largestPalindromes, findLargest(palindromes)...)
	}

	for _, v := range largestPalindromes {
		tmp, _ := strconv.Atoi(v)
		if largestPalindrome < tmp {
			largestPalindrome = tmp
		}
	}

	return nil, largestPalindrome

}
