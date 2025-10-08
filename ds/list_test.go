package ds

import (
    "container/list"
    "testing"
)

func TestList(t *testing.T) {
    l := list.New()
    for i := 0; i < 10; i++ {
        l.PushBack(i)
    }

    if l.Len() != 10 {
        t.Errorf("expected length 10, got %d", l.Len())
    }

    i := 0
    for e := l.Front(); e != nil; e = e.Next() {
        if e.Value.(int) != i {
            t.Errorf("expected value %d, got %d", i, e.Value)
        }
        i++
    }
}
