/* берем первое число и ищем его повтор
 * если повтор найден, записываем и берем второе число и ищем его повтор
 * иначе берем следующее число и повторяем цикл
 */
package task4

import (
	"strconv"
	"strings"
)
func numToSliceOfString(number int) []string {
	str := strconv.Itoa(number)
	sl := strings.Split(str, "")
	return sl
}

func isPal(sl []string) (b bool) {
	//
	return b
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func slToStr(sl []string) string {
	var str string
	for _, v := range sl {
		str += v
	}
	return str
}

func deleteElementFromSlice(s []string, index int) []string {
	newS := s[:index]
	newS = append(newS, s[index+1:]...)
	return newS
}

func foo(number int) (n int, ok bool) {
	sl := numToSliceOfString(number)
	for l, v := range sl {
		for r:=len(sl); r > l; r-- {
			if v == sl[r] && isPal(sl[l:r]) {
				return strToInt(slToStr(sl[l:r])), true
			}
		}
	}
	return 0, false
}


func DoTask4() {

}