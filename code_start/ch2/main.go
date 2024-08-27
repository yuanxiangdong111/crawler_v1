package main

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func main() {

    // firstMissingPositive
    // nums := []int{3, 4, -1, 19, -5}
    // positive := firstMissingPositive(nums)
    // fmt.Println(positive)

    // spiralOrder
    // matrix := [][]int{
    //     {1, 2, 3},
    //     {4, 5, 6},
    //     {7, 8, 9},
    // }
    // order := spiralOrder(matrix)
    // fmt.Println(order)

    // rotate
    // matrix := [][]int{
    //     {1, 2, 3, 4},
    //     {5, 6, 7, 8},
    //     {9, 10, 11, 12},
    //     {13, 14, 15, 16},
    // }
    // rotate(matrix)

    // 00
    // 01
    // 10
    // 11

    // searchMatrix
    matrix := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
        {13, 14, 15, 16},
    }
    b := searchMatrix(matrix, 7)
    fmt.Println(b)
}

func hasCycle(head *ListNode) bool {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }

    }
    return false
}

func isPalindrome(head *ListNode) bool {
    var nums []int

    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }

    for i := 0; i < len(nums)/2; i++ {
        if nums[i] != nums[len(nums)-1-i] {
            return false
        }
    }
    return true
}

func reverseList(head *ListNode) *ListNode {
    // 1 2 3 4 5
    // var pre *ListNode
    // cur := head
    //
    // for cur != nil {
    //     tmp := cur.Next
    //     cur.Next = pre
    //     pre = cur
    //     cur = tmp
    // }
    // return pre

    if head == nil || head.Next == nil {
        return head
    }

    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    fast, slow := headA, headB
    lenA, lenB := 0, 0
    for fast != nil {
        fast = fast.Next
        lenA++
    }
    for slow != nil {
        slow = slow.Next
        lenB++
    }

    var diffLen int
    if lenA > lenB {
        fast = headA
        slow = headB
        diffLen = lenA - lenB
    } else {
        slow = headA
        fast = headB
        diffLen = lenB - lenA
    }

    for i := 0; i < diffLen; i++ {
        fast = fast.Next
    }

    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }

    return fast
}

func searchMatrix(matrix [][]int, target int) bool {
    i, j := 0, len(matrix[0])-1

    for i < len(matrix) && j >= 0 {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] < target {
            i++
        } else {
            j--
        }
    }
    return false
}

func rotate(matrix [][]int) {
    // 1 2 3   7 4 1
    // 4 5 6   8 5 2
    // 7 8 9   9 6 3
    //  1  2  3  4     13  9  5  1
    //  5  6  7  8     14 10  6  2
    //  9 10 11 12     15 11  7  3
    // 13 14 15 16     16 12  8  4

    //  j,n-1-i   n-1-i,n-1-j  n-1-j,i
    // 13  9  5  1
    // 14 10  6  2
    // 15 11  7  3
    // 16 12  8  4

    // i,j  j,n-1-i

    n := len(matrix)
    for i := 0; i < n/2; i++ {
        for j := 0; j < (n+1)/2; j++ {
            tmp := matrix[i][j]
            matrix[i][j] = matrix[n-1-j][i]
            matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
            matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
            matrix[j][n-1-i] = tmp
        }
    }
    fmt.Println(matrix)
}

func spiralOrder(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
    var res []int
    tmp := make([][]bool, m)
    for i := 0; i < m; i++ {
        tmp[i] = make([]bool, n)
    }

    k := 1
    d := 0
    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}
    for i, j := 0, 0; k <= m*n; k++ {
        res = append(res, matrix[i][j])
        tmp[i][j] = true
        a := i + dx[d]
        b := j + dy[d]
        if a < 0 || a >= m || b < 0 || b >= n || tmp[a][b] == true {
            d = (d + 1) % 4
            a = i + dx[d]
            b = j + dy[d]
        }
        i, j = a, b

    }
    return res
}

func firstMissingPositive(nums []int) int {
    for _, num := range nums {
        if num <= 0 {
            num = len(nums) + 1
        }
    }

    for _, num := range nums {
        num = abs(num)
        if num <= len(nums) {
            nums[num-1] = -abs(nums[num-1])
        }
    }

    for index, _ := range nums {
        if nums[index] > 0 {
            return index + 1
        }
    }

    return len(nums) + 1

}

func abs(a int) int {
    if a <= 0 {
        return -a
    }
    return a
}
