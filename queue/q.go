package queue

type Queue struct {
	itens []any
}

func (q *Queue) Enqueue(item any) {
	q.itens = append(q.itens, item)
}

func (q *Queue) Dequeue() (any, bool) {
	if len(q.itens) <= 0 {
		return nil, true
	}

	//queue := []any(*q)
	lastItem := q.itens[0]
	q.itens = q.itens[1:]
	return lastItem, false
}
