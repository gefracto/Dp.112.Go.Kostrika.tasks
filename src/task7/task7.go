package task7

import (
	"os"
	"strconv"
	"strings"
)

func getFile(f string) string {

	file, err := os.Open(f)
	if err != nil {
		// handle the error here
		return f
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return f
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return f
	}

	str := string(bs)
	return str
}

func getLength(n uint64) int {
	str := strconv.Itoa(int(n))
	return len(str)
}

func Fibo(f string) []uint64 {
	var numbers []uint64
	if strings.Contains(f, " ") || strings.Contains(f, "\n") {
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
		// code
		length, _ := strconv.Atoi(f)
		//
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
