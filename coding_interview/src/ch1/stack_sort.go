package ch1stackqueue

type Stack struct {
    data []int
}

func (s *Stack) Push(v int) {
    s.data = append(s.data, v)
}

func (s *Stack) Pop() int {
    if len(s.data) == 0 {
        panic("Pop but data is empty")
    }

    v := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data)-1]

    return v
}

func (s *Stack) Top() int {
    if len(s.data) == 0 {
        panic("Pop but data is empty")
    }

    return s.data[len(s.data) - 1]
}

func (s *Stack) Size() int {
    return len(s.data)
}

func NewStack() *Stack {
    return &Stack{}
}

// Sort: min on top
func (s *Stack) Sort() {
    if len(s.data) <= 1 {
        return
    }

    helpStack := NewStack()
    for s.Size() != 0 {
        tmp := s.Pop()

        for helpStack.Size() > 0 && tmp < helpStack.Top() {
            s.Push(helpStack.Pop())
        }

        helpStack.Push(tmp)
    }

    for helpStack.Size() != 0 {
        s.Push(helpStack.Pop())
    }
}

func (s *Stack) IsSorted() bool {
    prev := -1
    for i, v := range s.data {
        if i > 0 && v > prev {
            return false
        }

        prev = v
    }

    return true
}

