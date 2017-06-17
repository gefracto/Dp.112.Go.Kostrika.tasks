package task6

import (
	"strconv"
	"math"
	"strings"
	"os"
)

func DoTask6(q int, sq int) {
	str := getSequence(q,sq)
	file, _ := os.Create("number sequence")
	defer file.Close()
	file.WriteString(str)
}

func getSequence(quantity int, square int) string {
	var file []string
	root := int(math.Floor(math.Sqrt(float64(square))))

	for root*root < square {
		root+=1
	}

	for i:=0; i < quantity; i++ {
		if root*root >= square {
			file = append(file, string(strconv.Itoa(root)))
		}
		root +=1
	}
	f := strings.Join(file, ", ")
	return f
}
