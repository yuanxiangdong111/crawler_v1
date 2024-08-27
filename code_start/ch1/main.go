package main

import (
    "fmt"
)

func main() {

    // maxSubArray
    // nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    // fmt.Println(maxSubArray(nums))

    // rotate
    // nums := []int{1, 2, 3, 4, 5, 6, 7}
    // rotate(nums, 3)

    // productExceptSelf
    nums := []int{2, 3, 4, 5}
    self := productExceptSelf(nums)
    fmt.Println(self)
}

func productExceptSelf(nums []int) []int {
    // nums = [2,3,4,5]
    // [60,40,30,24]
    // 60 20 5 1
    // 1  2  6 24
    n := len(nums)
    left := make([]int, n)
    left[0] = 1
    right := make([]int, n)
    right[n-1] = 1

    res := make([]int, n)
    for i := 1; i < n; i++ {
        left[i] = left[i-1] * nums[i-1]
    }

    for i := n - 2; i >= 0; i-- {
        right[i] = right[i+1] * nums[i+1]
    }

    for i := 0; i < n; i++ {
        res[i] = left[i] * right[i]
    }
    return res
}

func rotate(nums []int, k int) {
    // nums = [1,2,3,4,5,6,7], k = 3
    // [5,6,7,1,2,3,4]
    var tmp []int
    for _, num := range nums {
        tmp = append(tmp, num)
    }
    n := len(nums)
    for i := 0; i < n; i++ {
        nums[(i+k)%n] = tmp[i]
    }

}

func maxSubArray(nums []int) int {
    // nums = [-2,1,-3,4,-1,2,1,-5,4]
    // 6
    // [4,-1,2,1]

    // 贪心
    // tmp := 0
    // res := -200000000
    // for _, num := range nums {
    //     if tmp < 0 {
    //         tmp = num
    //     } else {
    //         tmp += num
    //     }
    //     res = getMaxNum(res, tmp)
    // }
    // return res

    // dp
    pre := 0
    res := nums[0]
    for _, num := range nums {
        pre = getMaxNum(num, pre+num)
        res = getMaxNum(res, pre)
    }
    return res
}

func getMaxNum(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 76
func minWindow(s string, t string) string {
    needStr := make(map[rune]int)

    for _, v := range t {
        needStr[v]++
    }

    startIndex := 0
    needCount := len(t)
    minLen := 20000000

    for l, r := 0, 0; r < len(s); r++ {
        if needStr[rune(s[r])] > 0 {
            needCount--
        }
        needStr[rune(s[r])]--

        if needCount <= 0 {
            for needStr[rune(s[l])] < 0 {
                needStr[rune(s[l])]++
                l++
            }

            if r-l+1 < minLen {
                minLen = r - l + 1
                startIndex = l
            }
            needStr[rune(s[l])]++
            needCount++
            l++
        }
    }
    return s[startIndex : startIndex+minLen]
}

func subarraySum(nums []int, k int) int {
    // numsLen := len(nums)
    // sum := 0
    // res := 0
    // for i := 0; i < numsLen; i++ {
    //
    //     if nums[i] == k {
    //         res++
    //     }
    //     sum = nums[i]
    //     for j := i + 1; j < numsLen; j++ {
    //         sum += nums[j]
    //         if sum == k {
    //             res++
    //         }
    //     }
    // }
    // return res

    mapNum := map[int]int{0: 1}
    numsLen := len(nums)
    res, preNum := 0, 0
    for i := 0; i < numsLen; i++ {
        preNum += nums[i]
        if mapNum[preNum-k] > 0 {
            res += mapNum[preNum-k]
        }

        mapNum[preNum]++
    }
    return res
}

func findAnagrams(s string, p string) []int {
    // s = "cbaebabacd", p = "abc"   [0,6]
    // s = "abab", p = "ab" [0,1,2]
    // s = "baa", p = "aa" [1]
    // mapStr := make(map[byte]int)
    var res []int
    var need [26]int

    for i := 0; i < len(p); i++ {
        need[p[i]-'a']++
    }

    for l, r := 0, 0; r < len(s); r++ {
        need[s[r]-'a']--
        for need[s[r]-'a'] < 0 {
            need[s[l]-'a']++
            l++
        }
        if r-l+1 == len(p) {
            res = append(res, l)
        }
    }

    return res
}

func lengthOfLongestSubstring(s string) int {
    // "abcabcbb" 3
    // "pwwkew" 3
    res := 0
    mapStr := make(map[byte]int)
    j := 0
    for i := 0; i < len(s); i++ {
        mapStr[s[i]]++

        for mapStr[s[i]] > 1 {
            mapStr[s[j]]--
            j++
        }
        res = maxNum(res, i-j+1)
    }

    return res
}

func trap(height []int) int {
    res := 0
    heightLen := len(height)
    maxLeft := make([]int, heightLen)
    maxRight := make([]int, heightLen)
    maxLeft[0] = height[0]
    maxRight[heightLen-1] = height[heightLen-1]

    for i := 1; i < heightLen; i++ {
        maxLeft[i] = maxNum(height[i], maxLeft[i-1])
    }
    // for _, v := range maxLeft {
    //     fmt.Println(v)
    // }

    for i := heightLen - 2; i >= 0; i-- {
        maxRight[i] = maxNum(height[i], maxRight[i+1])
    }

    for k, v := range height {
        res += minNum(maxLeft[k], maxRight[k]) - v
    }
    return res
}

func maxNum(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func minNum(a, b int) int {
    if a < b {
        return a
    }
    return b
}
