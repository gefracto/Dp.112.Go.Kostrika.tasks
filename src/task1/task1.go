package task1

import (
	"encoding/json"
	"errors"
	"github.com/gefracto/kostrika-go-tasks/src/tools"
	"strings"
)

type Data struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Symbol string `json:"symbol"`
}

func init() {
	tools.RememberMe(1, Dotask)
}

func Dotask(js []byte) (error, []byte) {

	var d Data
	json.Unmarshal(js, &d)

	// fmt.Println(d)
	// fmt.Println(js)

	w := d.Width
	h := d.Height
	s := d.Symbol

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
