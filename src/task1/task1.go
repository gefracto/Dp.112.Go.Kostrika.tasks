package task1

import "strings"

func makeRow(switcher int, symbol string, width int) string {
	row := ""
	var s string
	if switcher == 1 {
		s = symbol + " "
	} else {
		s = " " + symbol
	}

	if width%2 == 0 {
		for i := 1; i <= width/2; i++ {
			row = row + s
		}
	} else {
		for i := 1; i <= (width-1)/2; i++ {
			row = row + s
		}
		if switcher == 1 {
			row = row + symbol
		} else {
			row = row + " "
		}
	}

	return row
}

func ChessBoardBad(width int, height int, symbol string) string {
	switcher := 1
	board := ""
	for i := 1; i <= height; i++ {
		board = board + makeRow(switcher, symbol, width)
		board = board + "\n"
		switcher *= -1
	}

	return board
}

func ChessBoardNice(width int, height int, black string, white string) string {
	var blackSquare bool = true
	var board string = ""

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if blackSquare {
				board = board + black
			} else {
				board = board + white
			}
			blackSquare = !blackSquare
		}
		blackSquare = !blackSquare
		board = board + "\n"
	}
	return board
}

func ChessBoardVeryNice(width int, height int) string {
	var blackSquare bool = true
	var board string = ""
	var halfWidth int
	var extra bool = true
	if width%2 == 0 {
		halfWidth = width / 2
		extra = false
	} else {
		halfWidth = (width - 1) / 2
		extra = true
	}
	for i := 0; i < height; i++ {
		if extra == false {
			if blackSquare {
				board = board + strings.Repeat("* ", halfWidth) + "\n"
			} else {
				board = board + strings.Repeat(" *", halfWidth) + "\n"
			}
		} else {
			if blackSquare {
				board = board + strings.Repeat("* ", halfWidth) + "*\n"
			} else {
				board = board + strings.Repeat(" *", halfWidth) + " \n"
			}
		}
		blackSquare = !blackSquare
	}
	return board
}
