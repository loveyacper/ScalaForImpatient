package ch1stackqueue

type sliceStack struct {
    data [][]int
}

func (m *sliceStack) Size() int {
    return len(m.data)
}

func (m *sliceStack) Push(v int) {
    sv := make([]int, 1)
    sv[0] = v

    m.data = append(m.data, sv)
}

func (m *sliceStack) PushAtTop(v int) {
    if len(m.data) == 0 {
        panic("PushAtTop but data is empty")
    }

    m.data[len(m.data) - 1] = append(m.data[len(m.data) - 1], v)
}


func (m *sliceStack) Pop() []int {
    if len(m.data) == 0 {
        panic("Pop but data is empty")
    }

    sv := m.data[len(m.data) - 1]
    m.data = m.data[:len(m.data)-1]

    return sv
}

func (m *sliceStack) Top() []int {
    if len(m.data) == 0 {
        panic("Get but data is empty")
    }

    return m.data[len(m.data) - 1]
}

// 目标:记录每一个元素的左右两边距离最近且比它小的下标，数组元素可能重复
// 方法:用一个栈按照自底向上由小到大的顺序记录下标;
// 初始为空，直接将index压入
// 如果栈顶不为空，假设下标值是top，分三种情况：
// 1 index 对应的value比top对应的大，则index直接入栈，break
// 2 index 对应的value和top对应的相等，则修改top，append index，break
// 3 index 对应的value比top对应的小，则循环弹出top，因为栈中top下面的元素是它左边最近且比它小的下标，而index则是右边最近且比它小的下标
// 记录top对应的左右下标值，循环第3步

type IntPair struct {First, Second int}

func LeftRightMinValues(nums []int) []IntPair {
    if len(nums) == 0 {
        return []IntPair{}
    }
    res := make([]IntPair, len(nums))
    s := &sliceStack{}
    for i, v := range nums {
        needPush := true

        for s.Size() > 0 {
            top := s.Top()
            if v > nums[top[0]] {
                break
            } else if v == nums[top[0]] {
                s.PushAtTop(i)
                needPush = false
                break
            } else {
                s.Pop()
                for _, iv := range top {
                    left := -1
                    if s.Size() > 0 {
                        next := s.Top()
                        left = next[len(next)-1]
                    }
                    res[iv] = IntPair{left, i}
                }
            }
        }

        if needPush {
            s.Push(i)
        }
    }

    for s.Size() > 0 {
        top := s.Pop()
        for _, iv := range top {
            left := -1
            if s.Size() > 0 {
                next := s.Top()
                left = next[len(next)-1]
            }
            res[iv] = IntPair{left, -1}
        }
    }

    return res
}

