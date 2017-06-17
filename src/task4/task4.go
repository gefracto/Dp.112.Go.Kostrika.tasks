/* берем первое число и ищем его повтор
 * если повтор найден, записываем и берем второе число и ищем его повтор
 * иначе берем следующее число и повторяем цикл
 */
package task4

import (
	"strconv"
)

func Palindrom(number int) string {
	str := strconv.Itoa(number)
	return str
}
