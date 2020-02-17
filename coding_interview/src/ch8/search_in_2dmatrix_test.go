package ch8arraymatrix

import (
   "testing"
)


func TestReplaceBlank(t* testing.T) {
    matrix := [][]int{{1,3,5,7},
                      {10,11,16,20},
                      {23,30,34,50}}

    var tests = []struct {
        input int
        want bool
    } {
        {0, false},
        {10, true},
        {15, false},
        {20, true},
        {25, false},
        {30, true},
        {60, false},
    }

    for _, test := range tests {
        if got := SearchMatrix(matrix, test.input); got != test.want {
            t.Errorf("SearchMatrix(%q) = %v", test.input, got)
        }
    }
}

func BenchmarkSearchMatrix(b* testing.B) {
    matrix := [][]int{{1,3,5,7},
                      {10,11,16,20},
                      {23,30,34,50}}
    for i := 0; i < b.N; i++ {
        _ = SearchMatrix(matrix, i)
    }
}

func BenchmarkSearchMatrix2(b* testing.B) {
    matrix := [][]int{{1,3,5,7},
                      {10,11,16,20},
                      {23,30,34,50}}
    for i := 0; i < b.N; i++ {
        _ = SearchMatrix2(matrix, i)
    }
}

