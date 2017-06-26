package task1

import (
	"errors"
	"strings"
)

func (T *T1) Dotask1() (err error, data interface{}) {

	err, data = DoTask(T.Width, T.Height, T.Symbol)
	return
}

type T1 struct {
	Width, Height int
	Symbol        string
}

func makeRow(switcher int, symbol string, width int) string {
	var row string
	var ds string
	if switcher == 1 {
		ds = symbol + strings.Repeat(" ", len(symbol))
	} else {
		ds = strings.Repeat(" ", len(symbol)) + symbol
	}

	if width%2 == 0 {
		for i := 1; i <= width/2; i++ {
			row += ds
		}
	} else {
		for i := 1; i <= (width-1)/2; i++ {
			row += ds
		}
		if switcher == 1 {
			row += symbol
		} else {
			row += strings.Repeat(" ", len(symbol))
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

func DoTask(width int, height int, symbol string) (error, string) {
	var board string
	var err error

	if ok, err := validate(width, height); ok == false {
		return err, board
	}

	switcher := 1
	for i := 1; i <= height; i++ {
		board += makeRow(switcher, symbol, width) + "\n"
		switcher *= -1
	}

	return err, board

}
