package ch1stackqueue

type MinStack struct {
    data []int
    min []int
}

func (m *MinStack) Push(v int) {
    m.data = append(m.data, v)
    if len(m.min) == 0 || m.GetMin() >= v {
        m.min = append(m.min, v)
    }
}

func (m *MinStack) Pop() int {
    if len(m.data) == 0 {
        panic("Pop but data is empty")
    }

    v := m.data[len(m.data) - 1]
    m.data = m.data[:len(m.data)-1]

    if v == m.GetMin() {
        m.min = m.min[:len(m.min)-1]
    }

    return v
}

func (m *MinStack) GetMin() int {
    if len(m.min) == 0 {
        panic("GetMin but data is empty")
    }

    return m.min[len(m.min) - 1]
}

