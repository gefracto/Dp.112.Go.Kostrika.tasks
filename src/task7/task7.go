package task7

import (
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
	"regexp"
)

func getLength(n int) int {
	str := strconv.Itoa(n)
	return len(str)
}

func DoTask7(file string) {
	var f string
	var numbers []int

	if file != "context" {
		fmt.Println("Не удалось найти файл \"context\"")

	} else if contents, err := ioutil.ReadFile(file); err != nil {
		fmt.Println("Не удалось найти файл")

	} else if f = string(contents); len(f) == 0 {
		fmt.Println("Файл должен содержать числовые данные")

	} else if strings.Contains(f, " ") {
		s := strings.Split(f, " ")
		min, _ := strconv.Atoi(s[0])
		max, _ := strconv.Atoi(s[1])
		var a, b int = 1, 1
		for a < int(max) {
			a, b = b, a+b
			if a >= min && a < max {
				numbers = append(numbers, a)
			}
		}
		fmt.Println(numbers)

	} else if ok, _ := regexp.MatchString("([0-9]+)", f); ok {
		length, _ := strconv.Atoi(f)
		var a, b int = 1, 1
		for getLength(a) <= length {
			a, b = b, a+b
			if getLength(a) == length {
				numbers = append(numbers, a)

			}

		}
		fmt.Println(numbers)
	}

}
