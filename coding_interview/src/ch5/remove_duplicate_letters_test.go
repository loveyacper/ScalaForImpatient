package ch5string

import (
   "testing"
)

func TestRemoveDuplicate(t* testing.T) {
    if s := RemoveDuplicateLetters("cbacdcbc"); s != "acdb" {
        t.Errorf(`expect %v but %v`, "acdb", s)
    }
}

