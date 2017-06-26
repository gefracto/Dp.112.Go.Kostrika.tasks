package task1

import (
	"errors"
	"strings"
	//"fmt"
)

func (T *T1) Dotask1() (data string, err error){

	_, data, err = DoTask(T.Width, T.Height, T.Symbol)
	return data, err
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

func DoTask(width int, height int, symbol string) (ok bool, board string, err error) {
	if width <= 0 || height <= 0 {
		err := errors.New("Числа должны быть больше нуля!\n")
		return false, board, err
	} else {
		switcher := 1
		board := ""
		for i := 1; i <= height; i++ {
			board += makeRow(switcher, symbol, width) + "\n"
			switcher *= -1
		}
		return true, board, err
	}
}
