package task1

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
			row += s
		}
	} else {
		for i := 1; i <= (width-1)/2; i++ {
			row += s
		}
		if switcher == 1 {
			row += symbol
		} else {
			row += " "
		}
	}
	return row
}

func DoTask1(width uint, height uint, symbol string) string {
	switcher := 1
	board := ""
	for i := 1; i <= int(height); i++ {
		board += makeRow(switcher, symbol, int(width)) + "\n"
		switcher *= -1
	}
	return board
}
