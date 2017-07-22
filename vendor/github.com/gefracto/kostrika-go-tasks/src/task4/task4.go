package task4

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"

	"github.com/gefracto/kostrika-go-tasks/src/tasklist"
)

func init() {
	tasklist.RememberMe(4, Dotask4)
}

type T4 struct {
	Number int
}

func Dotask4(js []byte) (error, []byte) {
	var d T4
	json.Unmarshal(js, &d)
	err, in := dotask(d.Number)
	data := []byte(strconv.Itoa(in))
	return err, data
}

func reverseString(s string) string {
	var newString string = ""
	for i := len(s) - 1; i >= 0; i-- {
		newString += string(s[i])
	}
	return newString
}

func isPal(n []int) (b bool) {
	first := slicetonum(n) == slicetonum(reversenum(n))
	second := len(strconv.Itoa(slicetonum(n))) > 1
	if first && second {
		b = true
	}
	return b
}

func findLargest(sl []int) []int {
	var largest []int
	var largestSize int = 1
	for _, v := range sl {
		if len(strconv.Itoa(v)) > largestSize {
			largestSize = len(strconv.Itoa(v))
		}
	}
	for _, v := range sl {
		if len(strconv.Itoa(v)) == largestSize {
			largest = append(largest, v)
		}
	}
	return largest
}

func slicetonum(sl []int) int {
	l := float64(len(sl))
	n := 0.0
	pow := l
	for i := 0; i < int(l); i++ {
		n += float64(sl[i]) * math.Pow(10, pow)
		pow -= 1
	}
	return int(n) / 10
}
func numtoslice(num int) []int {
	l := len(strconv.Itoa(num))
	tmp := make([]int, l)

	for i := l - 1; i >= 0; i-- {
		tmp[i] = num % 10
		num = num / 10
	}
	return tmp
}

func reversenum(num []int) []int {
	var rev []int
	for i := len(num) - 1; i >= 0; i-- {
		rev = append(rev, num[i])
	}
	return rev
}

func dotask(num int) (error, int) {
	var ok bool
	ABC := numtoslice(num)
	CBA := reversenum(ABC)
	var largestPalindrome int
	var palindromes []int
	var largestPalindromes []int
	if num <= 0 {
		return errors.New("\nОшибка!\nЗадайте число больше нуля"), 0

	}

	for i, v := range ABC {
		for i2, v2 := range CBA {
			one := i+i2 < len(ABC)
			two := v2 == v
			if one && two && isPal(ABC[i:len(ABC)-i2]) {
				palindromes = append(palindromes, slicetonum(ABC[i:len(ABC)-i2]))

				ok = true
			}
		}
	}

	if ok {
		largestPalindromes = append(largestPalindromes, findLargest(palindromes)...)
	}

	for _, v := range largestPalindromes {
		if largestPalindrome < v {
			largestPalindrome = v
		}
	}

	return nil, largestPalindrome

}
