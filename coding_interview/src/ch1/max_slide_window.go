package ch1stackqueue

func MaxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    if k <= 0 {
        panic ("window size should > 0")
    }

    if k > len(nums) {
        panic ("window size should < len(nums)")
    }

    resLen := len(nums) - k + 1
    res := make([]int, 0, resLen)
    q := make([]int, 0, k)

    for i, _ := range nums {
        updateIndex(nums, i, k, &q)
        if i + 1 >= k {
            res = append(res, nums[q[0]])
        }
    }

    return res
}

// q 存放下标，按照对应的value从大到小的顺序
// 因此front的下标对应的值最大；
// 当遍历到新的下标index，假设值是vi, 只需要从头到尾比较，遇到了比
// vi大的元素，则继续遍历；
// 否则，vi更大或者相等但是vi下标更大，这样，q后面的元素就毫无意义
// 直接截断，并将vi下标index放在尾部即可
func updateIndex(nums []int, i int, k int, q* []int) {
    // 因为front的值更大，所以后面元素的值更小，但一定下标比front的要大
    // 也就是front的下标是最小的
    // 所以判断一下front是否过期了
    if len(*q) > 0 {
        frontIdx := (*q)[0]
        if frontIdx + k <= i {
            *q = (*q)[1:]
        }
    }

    // qv is index for nums
    for qi, qv := range *q {
        if nums[qv] <= nums[i] {
            *q = (*q)[:qi]
            break
        }
    }

    *q = append(*q, i)
}

