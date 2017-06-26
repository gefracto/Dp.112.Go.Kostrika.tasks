package task6

import (
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

type T6 struct {
	Length, MaxSquare int
}

func (T *T6) Dotask6() (err error, data interface{}) {
	return Dotask(T.Length, T.MaxSquare)
}

func Dotask(length, maxSquare int) (err error, b bool) {
	if length < 0 || maxSquare < 0 {
		return errors.New("Значения не должны быть меньше нуля.\nФайл не создан."), false

	} else {
		str := getSequence(length, maxSquare)
		file, err := os.Create("number sequence of task6")
		defer file.Close()
		_, err = file.WriteString(str)
		return err, true
	}
}

func getSequence(quantity, square int) string {
	var file []string
	root := int(math.Floor(math.Sqrt(float64(square))))

	for root*root < square {
		root += 1
	}

	for i := 0; i < quantity; i++ {
		if root*root >= square {
			file = append(file, string(strconv.Itoa(root)))
		}
		root += 1
	}

	f := strings.Join(file, ", ")
	return f
}
