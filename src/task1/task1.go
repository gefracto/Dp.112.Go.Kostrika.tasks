package task1

import (
	"errors"
	"github.com/gefracto/kostrika-go-tasks/src/tools"
	"github.com/json-iterator/go"
	"strings"
)

type Data struct {
	width, height int
	symbol        string
}

func init() {
	tools.RegisterJsonRunner(1, Dotask)
}

func Dotask(js []byte) (error, []byte) {

	var d Data
	jsoniter.Unmarshal(js, &d)

	w := d.width
	h := d.height
	s := d.symbol

	var board string
	var err error

	if ok, err := validate(w, h); ok == false {
		b := []byte(board)
		return err, b
	}

	switcher := true
	for i := 1; i <= h; i++ {
		board += makeRow(switcher, s, w) + "\n"
		switcher = !switcher
	}
	b := []byte(board)
	return err, b

}

func makeRow(switcher bool, s string, w int) string {
	var row string
	var ds string
	if switcher {
		ds = s + strings.Repeat(" ", len(s))
	} else {
		ds = strings.Repeat(" ", len(s)) + s
	}

	if w%2 == 0 {
		for i := 1; i <= w/2; i++ {
			row += ds
		}
	} else {
		for i := 1; i <= (w-1)/2; i++ {
			row += ds
		}
		if switcher {
			row += s
		} else {
			row += strings.Repeat(" ", len(s))
		}
	}
	return row
}

func validate(w, h int) (bool, error) {
	if w <= 0 || h <= 0 {
		return false, errors.New("Числа должны быть больше нуля!\n")
	}
	return true, nil
}
