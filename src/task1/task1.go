package task1

import (
	"errors"
	"strings"
)

func Dotask(w, h int, s string) (error, string) {
	var board string
	var err error

	if ok, err := validate(w, h); ok == false {
		return err, board
	}

	switcher := true
	for i := 1; i <= h; i++ {
		board += makeRow(switcher, s, w) + "\n"
		switcher = !switcher
	}

	return err, board

}

type T1 struct {
	Width, Height int
	Symbol        string
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

//func dotask(w int, h int, s string) (error, string) {
//	var board string
//	var err error

//	if ok, err := validate(w, h); ok == false {
//		return err, board
//	}

//	switcher := true
//	for i := 1; i <= h; i++ {
//		board += makeRow(switcher, s, w) + "\n"
//		switcher = !switcher
//	}

//	return err, board

//}
