package ch1stackqueue

import (
   "testing"
)

func TestCatDog(t* testing.T) {
    const (
        PUSH_DOG = iota
        PUSH_CAT

        POLL_DOG
        POLL_CAT
        POLL_ALL

        DOG_EMPTY
        CAT_EMPTY
        ALL_EMPTY
    )

    var tests = []struct {
        action int
        want_str string
        want_bool bool
    } {
        {PUSH_DOG, "", false},
        {PUSH_CAT, "", false},
        {PUSH_DOG, "", false},
        {PUSH_DOG, "", false},
        {PUSH_CAT, "", false},

        {POLL_ALL, "dog", false},
        {POLL_DOG, "dog", false},

        {POLL_ALL, "cat", false},

        {DOG_EMPTY, "", false},
        {CAT_EMPTY, "", false},

        {POLL_ALL, "dog", false},
        {POLL_ALL, "cat", false},

        {ALL_EMPTY, "", true},
    }

    q := NewCatDogQueue()
    for _, test := range tests {
        if test.action == PUSH_DOG {
            q.Add(Dog{})
        } else if test.action == PUSH_CAT {
            q.Add(Cat{})
        } else if test.action == POLL_DOG {
            if got := q.PollDog(); got.GetPetType() != test.want_str {
                t.Errorf("CatDogQueue pop (%v), want %v", got.GetPetType(), test.want_str)
            }
        } else if test.action == POLL_CAT{
            if got := q.PollCat(); got.GetPetType() != test.want_str {
                t.Errorf("CatDogQueue pop (%v), want %v", got.GetPetType(), test.want_str)
            }
        } else if test.action == POLL_ALL{
            if got := q.PollAll(); got.GetPetType() != test.want_str {
                t.Errorf("CatDogQueue pop (%v), want %v", got.GetPetType(), test.want_str)
            }
        } else if test.action == DOG_EMPTY {
            if got := q.IsDogEmpty(); got != test.want_bool {
                t.Errorf("CatDogQueue IsDogEmpty(%v), want %v", got, test.want_bool)
            }
        } else if test.action == CAT_EMPTY {
            if got := q.IsCatEmpty(); got != test.want_bool {
                t.Errorf("CatDogQueue IsCatEmpty(%v), want %v", got, test.want_bool)
            }
        } else if test.action == ALL_EMPTY {
            if got := q.IsEmpty(); got != test.want_bool {
                t.Errorf("CatDogQueue IsEmpty(%v), want %v", got, test.want_bool)
            }
        } else {
            t.Errorf("CatDogQueue unknown action (%q)", test.action)
        }
    }
}

