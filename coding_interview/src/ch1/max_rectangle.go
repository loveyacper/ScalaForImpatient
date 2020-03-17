package ch1stackqueue

// leetcode 84: 直方图的最大矩形
// 目标:计算数组组成的最大矩形面积:记录每一个元素的左右两边距离最近且比它小的下标，数组元素可能重复
// 原理:枚举每一个数字，以它为中心向左右扩展，遇到一个比它小的数则终止;如果左边没有比它小的，则记左下标-1；如果右边
// 没有比它小的，则记右下标为数组长度，面积就是val * (right - left - 1)
// 方法:用一个栈按照自底向上由小到大的顺序记录下标;
// 初始为空，直接将index压入
// 如果栈顶不为空，假设下标值是top，分2种情况：
// 1 index 对应的value比top对应的大，则index直接入栈，break
// 2 index 对应的value比top对应的小或相等，则循环弹出top，因为栈中top下面的元素是它左边最近且比它小的下标，而index则是右边最近且比它小的下标
// 记录top对应的左右下标值，循环第3步

func MaxRectangle(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    area := 0
    s := []int{} // as stack
    for i, v := range nums {
        for len(s) > 0 {
            top := s[len(s)-1]
            if v > nums[top] {
                break
            }

            s = s[:len(s)-1]
            right := i
            left := -1
            if len(s) > 0 {
                left = s[len(s)-1]
            }

            tmpArea := nums[top] * (right - left - 1)
            if tmpArea > area {
                area = tmpArea
            }
        }

        s = append(s, i)
    }

    right := len(nums)
    for len(s) > 0 {
        top := s[len(s)-1]
        s = s[:len(s)-1]
        left := -1
        if len(s) > 0 {
            left = s[len(s)-1]
        }

        tmpArea := nums[top] * (right - left - 1)
        if tmpArea > area {
            area = tmpArea
        }
    }

    return area
}

