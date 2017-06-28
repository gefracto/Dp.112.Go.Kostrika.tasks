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


func getLength(n int) int {
	return len(strconv.Itoa(n))
}


func validateFileName(s string) error {
	if s != "context" {
		return errors.New("Ошибка!\n" +
			"Не удалось найти файл.")
	}
	return nil

}

func canBeRead(s string) ([]byte, error) {
	contents, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, errors.New("Ошибка!\n"+
			"Не удалось прочитать файл.")
	}
	return contents, nil
}


func validateHasData(s string) error {
	if len(s) == 0 {
		return errors.New("Ошибка!\n"+
			"Файл должен содержать числовые данные")
	}
	return nil
}

func hasTwoNums(s string) bool {
	pattern := "\\A[\\d]+[,| |\n]{1}[\\d]+$"
	ok, _ := regexp.MatchString(pattern, s)
	return ok
	
}

func hasOneNum(s string) bool {
	pattern := "\\A[\\d]+$"
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

func splitData(s string) []string {
	if strings.Contains(s, "\n") {
		return strings.Split(s, "\n")

	}

	if strings.Contains(s, " ") {
		return strings.Split(s, " ")

	}

	if strings.Contains(s, ",") {
		return strings.Split(s, ",")
	}
	return nil
}

func validateNums(min, max int) (error, bool) {
	if min > max {
		return errors.New("Ошибка!\nЗначение нижней границы диапазона\n"+
			"не может быть больше значения верхней"), false
	}

	if min < 1 || max < 1 {
		return errors.New("Ошибка!\nЗначения должны быть больше нуля"), false
	}
	return nil, true
}

func caseTwoNums(FILE string) (err error, nums []int) {
	s := splitData(FILE)

	min, _ := strconv.Atoi(s[0])
	max, _ := strconv.Atoi(s[1])

	if err, ok := validateNums(min,max); ok == false {
		return err, nums
	}
	var a, b int = 1, 1
	for a < int(max) {
		a, b = b, a+b
		if a >= min && a < max {
			nums = append(nums, a)
		}
	}
	return err, nums
}

func caseOneNum(FILE string) (err error, nums[]int) {
	length, _ := strconv.Atoi(FILE)
	var a, b int = 1, 1

	for getLength(a) <= length {
		a, b = b, a+b
		if getLength(a) == length {
			nums = append(nums, a)

		}

	}
		return err, nums
}


func Dotask(file string) (err error, data []int) {
	var FILE string
	var numbers []int

	if validateFileName(file) != nil {
		return err, data
	}

	contents, err := canBeRead(file)
	if err != nil {
		return err, data

	}

	FILE = string(contents)

	if err := validateHasData(FILE); err != nil {
		return err, data

	}

	if hasTwoNums(FILE) {

		err, numbers = caseTwoNums(FILE)
		return err, numbers

	} else if hasOneNum(FILE) {

		err, numbers = caseOneNum(FILE)
		return err, numbers
	}

	return errors.New("Ошибка!\nФайл должен содержать либо:\n"+
		"1: Нижнюю и верхнюю границы " +
		"диапазона числовых значений,\n"+
		"разделенные запятой, пробелом или переносом строки." +
		"\n2: Числовое значение длины."), data

}
