package slice

import (
    "fmt"
    "testing"
)

func TestSlice(t *testing.T) {
    s1 := make([]int, 0, 10)
    // s1[:5] is a slice of length 0, because now s1 has length 0
    copy(s1[:5], []int{1, 2, 3, 4, 5})
    fmt.Println(s1)

    fmt.Println(s1)

    s2 := make([]int, 5, 10)
    copy(s2[:5], []int{1, 2, 3, 4, 5})
    s2[0] = 100
    fmt.Println(s2) // [100, 2, 3, 4, 5]

    // reallocation doesn't happen
    // s3 points to the same underlying array as s2
    s3 := append(s2, 6)
    s2[1] = 200
    fmt.Println(s3)

    // reallocation doesn't happen
    // s3 still points to the same underlying array
    s3 = append(s3, 7, 8, 9, 10)
    s2[2] = 300
    fmt.Println(s3)

    // capacity is exceeded. reallocation happens
    // s3 now points to a new underlying array
    s3 = append(s3, 11)
    s2[3] = 400
    fmt.Println(s3)
}

func TestSlice1(t *testing.T) {
    nums := []int{1, 2, 3, 4, 5}
    copy(nums[1:3], []int{8, 9}) // replaces elements at index 1 and 2
    fmt.Println(nums)            // [1 8 9 4 5]
}
