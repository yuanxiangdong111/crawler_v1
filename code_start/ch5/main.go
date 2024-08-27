package main

import (
    "bytes"
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func main() {
    // var stack []map[*TreeNode]int
    // pair := make(map[*TreeNode]int， root)
    // stack = append(stack, pair)
    // node2 := &TreeNode{Val: 3}
    // node1 := &TreeNode{Val: 2, Left: node2}
    // root := &TreeNode{Val: 1, Right: node1}

    // var root *TreeNode
    // var node1 *TreeNode
    // var node2 *TreeNode

    // inorderTraversal(root)
    //
    // fmt.Println("111")
    //
    // a := "dwdqwdqd"
    // b := "--yuanxiangdong"
    // strings := ConcatenateStrings(a, b)
    // fmt.Println("strings = ", strings)

    fmt.Println(lengthOfLongestSubstring("abcabcbb"))

    fmt.Println("dadsada2")
}

func isSymmetric(root *TreeNode) bool {

    // if root == nil{
    //     return true
    // }
    // var dfs func(l, r *TreeNode)bool
    //
    // dfs = func(l, r *TreeNode)bool{
    //     if l == nil && r == nil{
    //         return true
    //     }
    //     if l == nil || r == nil{
    //         return false
    //     }
    //
    //     if l.Val != r.Val{
    //         return false
    //     }
    //
    //
    //
    //     return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
    // }
    // return dfs(root.Left, root.Right)

    var stack []*TreeNode
    stack = append(stack, root.Left, root.Right)

    for len(stack) != 0 {
        node1, node2 := stack[0], stack[1]
        stack = stack[2:]
        if node1 == nil && node2 == nil {
            continue
        }

        if node1 == nil || node2 == nil || (node1.Val != node2.Val) {
            return false
        }

        stack = append(stack, node1.Left, node2.Right, node1.Right, node2.Left)
    }
    return true
}

func lengthOfLongestSubstring(s string) int {
    // abcabcbb 3
    strMap := make(map[byte]int)
    j := 0
    n := len(s)
    var res int
    for i := 0; i < n; i++ {
        strMap[s[i]]++
        fmt.Println(strMap)
        for strMap[s[i]] > 1 {
            strMap[s[j]]--
            j++
        }
        res = maxNum(res, i-j+1)
    }
    return res
}

func rightSideView(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    var stack []*TreeNode
    stack = append(stack, root)

    for len(stack) != 0 {
        size := len(stack)
        for i := 0; i < size; i++ {
            node := stack[0]
            stack = stack[1:]
            if node.Left != nil {
                stack = append(stack, node.Left)
            }
            if node.Right != nil {
                stack = append(stack, node.Right)
            }
            if i == size-1 {
                res = append(res, node.Val)
            }
        }
    }
    return res
}

// 有多条路径满足从根节点到叶子节点满足目标和
// 将满足的临时保存下来，在赋值到res结果中
func pathSum2(root *TreeNode, targetSum int) [][]int {
    var res [][]int
    var tmp []int
    var dfs func(root *TreeNode, curNum int)

    dfs = func(root *TreeNode, curNum int) {
        if root == nil {
            return
        }
        curNum += root.Val
        // 将每个节点中的值添加到tmp中
        tmp = append(tmp, root.Val)
        if curNum == targetSum && root.Left == nil && root.Right == nil {
            // 注意创建变量tt时，需要copy操作，且需要需要足够的大小
            // 如果直接 res = append(res, tmp) 这是用的引用，后续改变tmp的时候，res中的也会改变
            // tt := make([]int, len(tmp))
            // copy(tt, tmp)
            res = append(res, tmp)
        }

        // 递归向两边查询
        dfs(root.Left, curNum)
        dfs(root.Right, curNum)
        // 恢复现场
        tmp = tmp[0 : len(tmp)-1]
    }
    dfs(root, 0)
    return res
}

// 常规递归，如果是叶子节点且刚好满足从根节点到叶子节点是目标值的话，直接返回true (||的短路特性)
func hasPathSum(root *TreeNode, targetSum int) bool {

    var dfs func(root *TreeNode, curNum int) bool
    dfs = func(root *TreeNode, curNum int) bool {
        if root == nil {
            return false
        }
        curNum += root.Val
        if curNum == targetSum && root.Left == nil && root.Right == nil {
            return true
        }

        return dfs(root.Left, curNum) || dfs(root.Right, curNum)
    }

    return dfs(root, 0)
}

// 使用前缀和的思想，每次向左右两边递归的时候，记录下当前的累计和
// 如果当前的累积和减去目标值在其中一个前缀和中，表示找到了一条符合的路径，以此类推
func pathSum(root *TreeNode, targetSum int) int {
    preNum := make(map[int]int)
    // 判断如果根节点就是一个满足的路径
    preNum[0] = 1

    var dfs func(root *TreeNode, curNum int) int

    dfs = func(root *TreeNode, curNum int) int {
        // 递归的结束条件
        if root == nil {
            return 0
        }

        curNum += root.Val
        count := 0

        // 判断当前的路径中是否有一条符合的路径在前缀和中
        val, ok := preNum[curNum-targetSum]
        if ok {
            // 记录满足条件的路径数量
            count = val
        }

        // 记录累积和
        preNum[curNum]++
        // 左右递归查找
        count += dfs(root.Left, curNum)
        count += dfs(root.Right, curNum)
        // 回溯现场
        preNum[curNum]--

        return count
    }
    return dfs(root, 0)
}

func diameterOfBinaryTree(root *TreeNode) int {
    var res int
    var dfs func(root *TreeNode) int

    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        res = max(res, l+r)
        return max(l, r) + 1
    }
    return res
}
func maxNum(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func levelOrder(root *TreeNode) [][]int {
    var res [][]int
    var queue []*TreeNode
    queue = append(queue, root)

    for len(queue) != 0 {
        var currentNodes []int
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            currentNodes = append(currentNodes, node.Val)
        }
        res = append(res, currentNodes)
    }

    return res
}

func ConcatenateStrings(s ...string) string {
    if len(s) == 0 {
        return ""
    }
    var buffer bytes.Buffer
    for _, i := range s {
        buffer.WriteString(i)
    }
    return buffer.String()
}

func inorderTraversal(root *TreeNode) []int {
    type Pair struct {
        Val  int
        Node *TreeNode
    }
    var stack []Pair
    var res []int
    nodeRoot := Pair{
        Val:  0,
        Node: root,
    }
    stack = append(stack, nodeRoot)

    for len(stack) != 0 {
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if cur.Node == nil {
            continue
        }

        if cur.Val == 0 {
            stack = append(stack, Pair{Val: 0, Node: cur.Node.Right})
            stack = append(stack, Pair{Val: 1, Node: cur.Node})
            stack = append(stack, Pair{Val: 0, Node: cur.Node.Left})
        } else {
            res = append(res, cur.Node.Val)
        }

    }
    return res
}
