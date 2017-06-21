package task6

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func DoTask6(length, maxSquare int) (ok bool, err error, reason string) {
	if length < 0 || maxSquare < 0 {
		return false, err, "Значения не должны быть меньше нуля.\nФайл не создан."
	} else {
		str := getSequence(length, maxSquare)
		file, err := os.Create("number sequence of task6")
		defer file.Close()
		_, err = file.WriteString(str)
		return true, err, reason
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
