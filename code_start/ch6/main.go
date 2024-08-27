package main

import (
    "bytes"
    "fmt"
    "math"
    "sync"
)

var wg sync.WaitGroup
var ch1 chan int
var ch2 chan int
var ch3 chan int

func PA() {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        <-ch1
        fmt.Println("A")
        ch2 <- 1
    }
}

func PB() {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        <-ch2
        fmt.Println("B")
        ch3 <- 1
    }
}

func PC() {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        <-ch3
        fmt.Println("C")
        ch1 <- 1
    }
}

func main() {
    // Login = make(chan int, 1)
    // ch2 = make(chan int, 1)
    // ch3 = make(chan int, 1)
    // Login <- 1
    // wg.Add(1)
    // go PA()
    // go PB()
    // go PC()
    //
    // wg.Wait()
    // height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    // fmt.Println(trap(height))

    f := squares()
    fmt.Println(f())

    max := Max(1, 2, 3, 4, 5)
    fmt.Println("maxNum = ", max)
    min := Min(1, 2, 3, 4, 5)
    fmt.Println("minNum = ", min)
    str := JoinStr("aaa", "+sdas", "+2321")
    fmt.Println("str = ", str)
}

func A(a int) int {
    return 0
}

func A(a, b int) int {
    return 0
}

// 练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
func Max(nums ...int) int {
    if len(nums) == 0 {
        return math.MaxInt
    }
    i, n := 0, len(nums)
    for j := 0; j < n; j++ {
        if nums[j] > nums[i] {
            i = j
        }
    }
    return nums[i]
}

func Min(nums ...int) int {
    if len(nums) == 0 {
        return math.MinInt
    }
    i, n := 0, len(nums)
    for j := 0; j < n; j++ {
        if nums[j] < nums[i] {
            i = j
        }
    }
    return nums[i]
}

func JoinStr(str ...string) string {
    if len(str) == 0 {
        return ""
    }
    var buffer bytes.Buffer
    for _, v := range str {
        buffer.WriteString(v)
    }
    return buffer.String()
}

func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func dailyTemperatures(temperatures []int) []int {
    //      输入: temperatures = [73,74,75,71,69,72,76,73]
    //      输出: [1,1,4,2,1,1,0,0]
    // 维护一个单调递减的队列 当前的值大于栈尾值的时候，说明找到了第一个满足的数（即下一个升温的天）
    // 当前位置温度到下一个第一次升温天 所需要的天数 表示为res[index] = i - index
    // 其中index表示单调递减队列的队尾索引 i表示当前天数的索引
    var queue []int
    n := len(temperatures)
    res := make([]int, n)

    for i := 0; i < n; i++ {
        for len(queue) > 0 && temperatures[queue[len(queue)-1]] < temperatures[i] {
            index := queue[len(queue)-1]
            queue = queue[:len(queue)-1]
            res[index] = i - index
        }

        queue = append(queue, i)
    }

    return res
}

func trap(height []int) int {
    /*
       // 单调栈的做法
       // 维护一个单调递减的队列 直到遇到一个比队尾大的数
       // 记当前元素为 height[i] 队尾元素为queue[len(queue)-1]
       // 此时队尾元素所在的凹槽能积的雨水 取决于队尾元素的上一个元素和当前元素的最小值 减去 凹槽位置的高度 再乘以宽度
       // height = [0,1,0,2,1,0,1,3,2,1,2,1]
       var queue []int
       var res int

       for i := 0; i < len(height); i++ {
            // 维护单调队列直到遇到比队尾大的数
           for len(queue) > 0 && height[queue[len(queue)-1]] < height[i] {
            // 取出队尾元素高度
               h := height[queue[len(queue)-1]]
            // 出队
               queue = queue[:len(queue)-1]
            // 队列没元素了，表示遍历完了
               if len(queue) == 0 {
                   break
               }
            // 计算凹槽的雨水
               diff := GetMinNum(height[queue[len(queue)-1]], height[i])
               distance := i - queue[len(queue)-1] - 1
               res += distance * (diff - h)
           }
            // 单调队列入队
           queue = append(queue, i)
       }

       return res
    */

    // 动态规划
    n := len(height)
    leftMax := make([]int, n)
    rightMax := make([]int, n)
    var res int

    leftMax[0] = height[0]
    rightMax[n-1] = height[n-1]

    // 获取当前及其左边位置的最大值
    for i := 1; i < n; i++ {
        leftMax[i] = GetMaxNum(height[i], leftMax[i-1])
    }

    // 获取当前及其右边位置的最大值
    for i := n - 2; i >= 0; i-- {
        rightMax[i] = GetMaxNum(height[i], rightMax[i-1])
    }

    // 当前位置可以接的雨水取决于 当前位置左右两边柱子最低位置及其当前位置的高度
    for i := 1; i < n-1; i++ {
        res += GetMinNum(leftMax[i], rightMax[i]) - height[i]
    }
    return res
}

func GetMinNum(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func GetMaxNum(a, b int) int {
    if a > b {
        return a
    }
    return b
}
