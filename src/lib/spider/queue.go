package spider

//队列
type Queue struct {
	StacksA map[string]string
	StacksB map[string]string
}

//插入队列
func (que *Queue) QueueInsert(data map[string]string) {
	for key, val := range data {
		if _, ok := que.StacksB[key]; !ok {
			que.StacksA[key] = val
		}
	}
}

//从队列取出值
func (que *Queue) QueueShift() string {
	for key, val := range que.StacksA {
		que.StacksB[key] = val
		delete(que.StacksA, key)
		return val
	}
	return ""
}
