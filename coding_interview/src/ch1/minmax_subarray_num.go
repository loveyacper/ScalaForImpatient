package ch1stackqueue

// 目标:计算所有满足最大值与最小值之差<=gap的子数组数量
// 原理:首先需要能够快速找到子数组的最大最小值，使用两个队列来维护,O(1)
// 初始化两个变量表示子数组的边界.当子数组满足需求时，扩张右边界，直到不满足.
// 此时，计算以左边界为起点的子数组数量: right - left
// 然后递增左边界。继续循环
// 另外，在遍历过程中，别忘了维护队列，以查询最大最小值

func SubarrayNum(nums []int, gap int) int {
    if gap < 0 || len(nums) == 0 {
        return 0
    } else if gap == 0 {
        return len(nums)
    }

    count := 0
    qmax := []int{} // as queue, the max value's index is at front
    qmin := []int{} // as queue, the min value's index is at front
    start := 0
    end := 0 // [start, end] is a subarray
    for end < len(nums) {
        if len(qmax) == 0 {
            qmax = append(qmax, end)
            qmin = append(qmin, end)
            end++
        } else {
            max := nums[qmax[0]]
            if max < nums[end] {
                max = nums[end]
            }

            min := nums[qmin[0]]
            if min > nums[end] {
                min = nums[end]
            }

            if max - min <= gap {
                // enqueue
                updateAdd(nums, end, &qmin, &qmax)
                end++
            } else {
                // 统计以start为起点的满足条件的子数组数量
                count += (end - start)
                updateRemove(start, &qmin, &qmax)
                start++
                // dequeue
            }
        }
    }

    count += (end - start) * (1 + end - start) / 2 // equal diff sequence
    return count
}

func updateAdd(nums []int, add int, qmin* []int, qmax* []int) {
    v := nums[add]

    for i, iv := range *qmin {
        if nums[iv] >= v {
            *qmin = (*qmin)[:i]
            break
        }
    }
    *qmin = append(*qmin, add)

    for i, iv := range *qmax {
        if nums[iv] <= v {
            *qmax = (*qmax)[:i]
            break
        }
    }

    *qmax = append(*qmax, add)
}
func updateRemove(remove int, qmin* []int, qmax* []int) {
    if (*qmin)[0] == remove {
        (*qmin) = (*qmin)[1:]
    }

    if (*qmax)[0] == remove {
        (*qmax) = (*qmax)[1:]
    }
}

