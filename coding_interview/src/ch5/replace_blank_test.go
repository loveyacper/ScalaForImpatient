package ch5string

import (
   "testing"
)

func TestReplaceBlank(t* testing.T) {
    str := make([]byte, 0, 16)
    str = append(str, 'a', ' ', 'b', ' ')
    // 类似c语言的(void)unused
    _ = ReplaceBlank(&str)

    if len(str) != 8 {
        t.Errorf(`len should be 8 but %d`, len(str))
    }
}

func BenchmarkReplaceBlank(b* testing.B) {
    for i := 0; i < b.N; i++ {
        str := make([]byte, 0, 16)
        str = append(str, 'a', ' ', 'b', ' ')
        _ = ReplaceBlank(&str)
    }
}

