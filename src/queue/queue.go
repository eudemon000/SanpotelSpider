package queue

import (
	"container/list"
)

//队列
type msgQueen struct {
	list	*list.List		//队列
	count	int				//每次取出多少队列内容
	ch chan int
}

var q *msgQueen

func init() {
	q = new(msgQueen)
	q.list = list.New()
}

//向队列插入数据
func Push(data interface{}) {
	q.list.PushBack(data)
}

//从队列读取数据
func Pull(max int) []interface{} {
	//q.ch <- 0
	num := q.list.Len()
	if max < num {	//如果当前获取的数据大于队列的长度，则获取最大数为队列的长度
		num = max
	}
	q.count = num

	i := num
	l := make([]interface{}, 0)
	for v := q.list.Front(); v != nil && i > 0; v = v.Next() {
		if i > 0 {
			value := v.Value
			l = append(l, value)
			q.list.Remove(v)
		}
		i--
	}
	return l
}