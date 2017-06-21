package task1

import (
	"strings"
)

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

func DoTask1(width int, height int, symbol string) (ok bool, data interface{}, reason string) {
	if width <= 0 || height <= 0 {
		return false, nil, "Ошибка!\nЗначения ширины и высоты должно быть больше ноля\n"
	} else {
		switcher := 1
		board := ""
		for i := 1; i <= height; i++ {
			board += makeRow(switcher, symbol, width) + "\n"
			switcher *= -1
		}
		return true, board, reason
	}
}
