package task7

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type T7 struct {
	File string
}

func (T *T7) Dotask7() (err error, data interface{}) {
	return Dotask(T.File)
}

// разбить код из dotask на отдельные куски

func getLength(n int) int {
	str := strconv.Itoa(n)
	return len(str)
}

func Dotask(file string) (err error, data []int) {
	var f string
	var numbers []int
	var s []string

	if file != "context" {
		return errors.New("Ошибка!\nНе удалось найти файл."), data

	} else if contents, err := ioutil.ReadFile(file); err != nil {
		return errors.New("Ошибка!\nНе удалось прочитать файл."), data

	} else if f = string(contents); len(f) == 0 {
		return errors.New("Ошибка!\nФайл должен содержать числовые данные"), data

	} else if ok, _ := regexp.MatchString("\\A[\\d]+[,| |\n]{1}[\\d]+$", f); ok {

		if strings.Contains(f, "\n") {
			s = strings.Split(f, "\n")

		} else if strings.Contains(f, " ") {
			s = strings.Split(f, " ")

		} else if strings.Contains(f, ",") {
			s = strings.Split(f, ",")
		}

		min, _ := strconv.Atoi(s[0])
		max, _ := strconv.Atoi(s[1])

		if min > max {
			return errors.New("Ошибка!\nЗначение нижней границы диапазона\nне может быть больше значения верхней"), data

		} else if min < 1 || max < 1 {
			return errors.New("Ошибка!\nЗначения должны быть больше нуля"), data
		} else {
			var a, b int = 1, 1
			for a < int(max) {
				a, b = b, a+b
				if a >= min && a < max {
					numbers = append(numbers, a)
				}
			}

			return err, numbers
		}

	} else if ok, _ := regexp.MatchString("\\A[\\d]+$", f); ok {
		length, _ := strconv.Atoi(f)
		var a, b int = 1, 1

		for getLength(a) <= length {
			a, b = b, a+b
			if getLength(a) == length {
				numbers = append(numbers, a)

			}

		}

		return err, numbers
	}
	return errors.New("Ошибка!\nФайл должен содержать либо:\n1: Нижнюю и верхнюю границы " +
		"диапазона числовых значений,\nразделенные запятой, пробелом или переносом строки." +
		"\n2: Числовое значение длины."), data

}
