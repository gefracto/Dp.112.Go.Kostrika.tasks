package tasklist

type task func([]byte) (error, []byte)

var TasksMap map[int]task = make(map[int]task)

func RememberMe(tasknum int, t task) {
	TasksMap[tasknum] = t
	//	panic(t)
}
