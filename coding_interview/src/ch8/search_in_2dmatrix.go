package ch8arraymatrix

func SearchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    // 找右上角元素，如果比target小，第一行被排除；
    // 如果大，最后一列被排除

    col := len(matrix[0])

    pivot := matrix[0][col-1]
    if pivot == target {
        return true
    } else if pivot < target {
        //fmt.Printf("pivot %d < target %d\n", pivot, target)
        return SearchMatrix(matrix[1:][:], target)
    } else {
        //fmt.Printf("pivot %d > target %d\n", pivot, target)
        mtx := make([][]int, len(matrix))
        for i := 0; i < len(mtx); i++ {
            mtx[i] = matrix[i][:col-1]
        }

        return SearchMatrix(mtx, target)
        // 这样切片无法通过
        //return searchMatrix(matrix[:][:col-1], target)
    }
}

func SearchMatrix2(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    // 找右上角元素，如果比target小，第一行被排除；
    // 如果大，最后一列被排除

    row := 0
    col := len(matrix[0]) - 1

    for row < len(matrix) && col >= 0 {
        pivot := matrix[row][col]
        if pivot == target {
            return true
        } else if pivot < target {
            row++
        } else {
            col--
        }
    }

    return false
}

