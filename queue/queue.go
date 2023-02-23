package queue

type Queue[T any] struct {
	queue []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{make([]T, 0)}
}

func (q *Queue[T]) Enqueue(e T) {
	q.queue = append(q.queue, e)
}

func (q *Queue[T]) Dequeue() T {
	e := q.queue[0]

	if len(q.queue) == 0 {
		q.queue = make([]T, 1)
	} else {
		q.queue = q.queue[1:]
	}
	return e
}

func (q *Queue[T]) IsNotEmpty() bool {
	return len(q.queue) != 0
}
