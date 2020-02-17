package ch5string

//import "fmt"

// convert space to %20

func ReplaceBlank(str* []byte) string {
    //fmt.Printf("[%v] len %v, cap %v\n", str, len(*str), cap(*str))
    if str == nil || len(*str) == 0 {
        return string(*str)
    }

    spaces := 0
    for _, v := range *str {
        if v == ' ' {
            spaces += 1
        }
    }

    if spaces == 0 {
        return string(*str)
    }

    olen := len(*str)
    nlen := olen + spaces * 2
    if cap(*str) < nlen {
        panic("Not enough room")
    }

    // 相当于cpp的resize()
    for i := 0; i < (2 * spaces); i++ {
        *str = append(*str, byte(0))
    }

    for i, j := nlen-1, olen-1; j>=0 && i>=0; j-- {
        if (*str)[j] != ' ' {
            (*str)[i] = (*str)[j]
            i--
        } else {
            // convert space to %20
            (*str)[i] = '0'
            (*str)[i-1] = '2'
            (*str)[i-2] = '%'
            i = i - 3
        }
    }

    return string(*str)
}

