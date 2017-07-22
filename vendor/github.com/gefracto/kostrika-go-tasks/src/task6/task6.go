package task6

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gefracto/kostrika-go-tasks/src/tasklist"
)

func init() {
	tasklist.RememberMe(6, Dotask6)
}

type T6 struct {
	Length, MaxSquare int
}

func Dotask6(js []byte) (error, []byte) {
	var d T6
	var err error
	if err = json.Unmarshal(js, &d); err != nil {
		return err, []byte("Ошибка unmarshal")
	}
	err = dotask(d.Length, d.MaxSquare)
	return err, []byte(fmt.Sprint(err))
}

func dotask(length, maxSquare int) error {
	if length < 0 || maxSquare < 0 {
		return errors.New("Значения не должны быть меньше нуля.\nФайл не создан.")

	}
	str := getSequence(length, maxSquare)
	file, err := os.Create("number sequence of task6")
	defer file.Close()
	_, err = file.WriteString(str)
	return err

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

	str := strings.Join(file, ", ")
	return str
}
