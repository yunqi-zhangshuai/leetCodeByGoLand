package queue

type ArrayCircularQueue struct {
	qArr  [qSize]interface{}
	front int // 队首
	rear  int // 队尾
}

const qSize int = 100

func NewArrayCircularQueue() *ArrayCircularQueue {
	return &ArrayCircularQueue{}
}

func (cq *ArrayCircularQueue) EnQueue(item interface{}) {

}

func (cq *ArrayCircularQueue) DeQueue() (interface{}, error) {

	return nil, nil
}

func (cq *ArrayCircularQueue) isFull() bool {

	return false
}
