package task7

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func getLength(n int) int {
	str := strconv.Itoa(n)
	return len(str)
}

func DoTask7(file string) (ok bool, data []int, reason string) {
	var f string
	var numbers []int
	var s []string

	if file != "context" {
		return false, data, "Ошибка!\nНе удалось найти файл."

	} else if contents, err := ioutil.ReadFile(file); err != nil {
		return false, data, "Ошибка!\nНе удалось прочитать файл."

	} else if f = string(contents); len(f) == 0 {
		return false, data, "Ошибка!\nФайл должен содержать числовые данные"

	} else if ok, _ := regexp.MatchString("[\\d]+[,| |\n]{1}[\\d]+", f); ok {

		if  strings.Contains(f, "\n") {
			s = strings.Split(f, "\n")
		} else if strings.Contains(f, " ") {
			s = strings.Split(f, " ")
		} else if strings.Contains(f, ",") {
			s = strings.Split(f, ",")
		}

		min, _ := strconv.Atoi(s[0])
		max, _ := strconv.Atoi(s[1])

		if min > max {
			return false, data, "Ошибка!\nЗначение нижней границы диапазона\nне может быть больше значения верхней"

		} else if min < 1 || max < 1 {
			return false, data, "Ошибка!\nЗначения должны быть больше нуля"
		} else {
			var a, b int = 1, 1
			for a < int(max) {
				a, b = b, a+b
				if a >= min && a < max {
					numbers = append(numbers, a)
				}
			}

			return true, numbers, reason
		}

	} else if ok, _ := regexp.MatchString("[\\d]+", f); ok {
		length, _ := strconv.Atoi(f)
		var a, b int = 1, 1

		for getLength(a) <= length {
			a, b = b, a+b
			if getLength(a) == length {
				numbers = append(numbers, a)

			}

		}

		return true, numbers, reason
	}
	return false, data, "Ошибка!\nФайл должен содержать либо:\n1: Нижнюю и верхнюю границы " +
		"диапазона числовых значений,\nразделенные запятой, пробелом или переносом строки." +
		"\n2: Числовое значение длины."

}
