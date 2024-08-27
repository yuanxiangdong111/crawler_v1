package main

import (
    "bufio"
    "bytes"
    "fmt"
)

// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。

type WordCount int
type LinesCount int

func main() {
    // fmt.Println(NonZero())

    var w WordCount
    n, _ := w.Write([]byte("hello world aa"))
    fmt.Println(n)
    fmt.Println("w = ", w)
    w = 0

    // fmt.Fprintf第一个参数是io.Writer
    fmt.Fprintf(&w, "words count is %s", "hello")
    fmt.Println(w)
}

func NonZero() (num int) {

    // panic被恢复，设置返回值
    defer func() {
        if p := recover(); p != nil {
            num = 2
        }
    }()

    // 发生panic 跳到defer中
    panic("?")
}

// 计算行数
func (l *LinesCount) Write(p []byte) (line int, err error) {
    scanner := bufio.NewScanner(bytes.NewReader(p))
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        *l++
    }
    return int(*l), nil
}

func (w *WordCount) Write(p []byte) (n int, err error) {
    scanner := bufio.NewScanner(bytes.NewReader(p))
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        *w++
    }
    return int(*w), nil
}

func nextGreaterElements(nums []int) []int {
    var queue []int
    var res, newNums []int
    n := len(nums)
    mapNum := make(map[int]int)
    tmp := nums[:len(nums)-1]
    newNums = append(nums, tmp...)

    // 5 4 3 2 1 5 4 3 2
    // 1 5
    // 2 5
    // 3 5
    // 4 5
    for i := 0; i < len(newNums); i++ {

        for len(queue) > 0 && newNums[queue[len(queue)-1]] < newNums[i] {
            cur := newNums[queue[len(queue)-1]]
            queue = queue[:len(queue)-1]

            if mapNum[cur] <= 0 {
                mapNum[cur] = newNums[i]
            }

        }

        queue = append(queue, i)
    }

    for i := 0; i < n; i++ {
        if _, ok := mapNum[nums[i]]; ok {
            res = append(res, mapNum[nums[i]])
        } else {
            res = append(res, -1)
        }
    }

    return res
}
