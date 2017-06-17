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

func deleteElementFromSlice(s []string, index int) []string {
	newS := s[:index]
	newS = append(newS, s[index+1:]...)
	return newS
}

func foo(number int) (n int, ok bool) {
	sl := numToSliceOfString(number)

}


func DoTask4() {

}