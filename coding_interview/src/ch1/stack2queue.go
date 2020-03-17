package ch1stackqueue


type stack struct {
    data []int
}

func (s *stack) Push(v int) {
    s.data = append(s.data, v)
}

func (s *stack) Pop() int {
    if len(s.data) == 0 {
        panic("Pop but data is empty")
    }

    v := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data)-1]

    return v
}

func (s *stack) Top() int {
    if len(s.data) == 0 {
        panic("Pop but data is empty")
    }

    return s.data[len(s.data) - 1]
}

func (s *stack) Size() int {
    return len(s.data)
}


type Queue struct {
    head stack
    tail stack
}
// push pop Front

func (q *Queue) Push(v int) {
    q.tail.Push(v)
}

func (q *Queue) Pop() int {
    v := q.Front()
    _ = q.head.Pop()
    return v
}

func (q *Queue) Front() int {
    if q.head.Size() == 0 && q.tail.Size() == 0 {
        panic("Queue empty")
    }

    if q.head.Size() == 0 {
        for q.tail.Size() > 0 {
            v := q.tail.Pop()
            q.head.Push(v)
        }
    }

    return q.head.Top()
}

func NewQueue() *Queue {
    s1 := stack{}
    s2 := stack{}
    return &Queue{ head : s1, tail : s2 }
}

