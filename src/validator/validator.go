package validator

import (
	jsn "jsonToStruct"
	t1 "task1"
	t2 "task2"
	t3 "task3"
	t4 "task4"
	t5 "task5"
	t6 "task6"
	t7 "task7"
	"fmt"
)
func separate(n int) {
	fmt.Printf("\nTask #%d\n----------------------\n", n)
}

func Dance(f string){

	separate(1)
	fmt.Println(t1.DoTask1(5, 8, "*"))

	separate(2)
	fmt.Println(t2.DoTask2(t2.Envelope{21, 15}, t2.Envelope{3.5, 22.3}))

	separate(3)
	s := make([]t3.Triangle, 1)
	s = append(s, t3.Triangle{"ABC", 10.3, 10.1, 10}, t3.Triangle{"DEF", 10.1, 10.3, 10}, t3.Triangle{"GHI", 34, 43.2, 20}, t3.Triangle{"JKL", 2, 3, 2})
	fmt.Println(t3.DoTask3(s))

	separate(4)
	fmt.Println(t4.DoTask4(41212226))
	fmt.Println(t4.DoTask4(1))
	fmt.Println(t4.DoTask4(43956))

	separate(5)
	fmt.Println(t5.DoTask5(000001, 999999))

	separate(6)
	fmt.Println(t6.DoTask6(10, 25))

	separate(7)
	fmt.Println(t7.DoTask7("context"))  // 1 100
	fmt.Println(t7.DoTask7("context2")) // 2
}
