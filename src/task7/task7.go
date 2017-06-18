package task7

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func getFile(f string) (string, bool) {
	ok := false
	contents, _ := ioutil.ReadFile(f)
	str := string(contents)
	return str, ok
}

func getLength(n uint64) int {
	str := strconv.Itoa(int(n))
	return len(str)
}

func DoTask7(file string) []uint64 {
	f, _ := getFile(file)
	var numbers []uint64

	if strings.Contains(f, " ") {
		s := strings.Split(f, " ")
		min, _ := strconv.Atoi(s[0])
		max, _ := strconv.Atoi(s[1])
		var a, b uint64 = 1, 1
		for a < uint64(max) {
			a, b = b, a+b
			if a >= uint64(min) && a < uint64(max) {
				numbers = append(numbers, a)
			}
		}
	} else {
		length, _ := strconv.Atoi(f)
		var a, b uint64 = 1, 1
		for getLength(a) <= length {
			a, b = b, a+b
			if getLength(a) == length {
				numbers = append(numbers, a)

			}

		}
	}
	return numbers

}
