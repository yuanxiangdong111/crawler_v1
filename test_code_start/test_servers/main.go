package main

import (
    "fmt"
)

func main() {
    var a = []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 6, 7, 8, 9}

    // 二分查找模版,左边第一个满足的数
    l, r := 0, len(a)-1
    for l < r {
        mid := (l + r) / 2
        if 5 <= a[mid] {
            r = mid
        } else {
            l = mid + 1
        }
    }

    fmt.Printf("index : %d, value : %d\n", l, a[l])

    // 二分查找模版,右边第一个满足的数
    l, r = 0, len(a)-1
    for l < r {
        mid := (l + r + 1) / 2
        if a[mid] <= 5 {
            l = mid
        } else {
            r = mid - 1
        }
    }
    fmt.Printf("index : %d, value : %d\n", l, a[l])

    a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    l, r = 0, len(a)
    mid := (r + l) / 2
    for l+1 < r {
        if 5 < a[mid] {
            r = mid
        } else {
            l = mid
        }
        mid = (r + l) / 2
    }

    fmt.Printf("index : %d, value : %d\n", l, a[l])
}
