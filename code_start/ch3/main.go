package main

type ListNode struct {
    Val  int
    Next *ListNode
}

type Node struct {
    Val    int
    Next   *Node
    Random *Node
}

func main() {

}

func sortList(head *ListNode) *ListNode {
    return ListSort(head, nil)
}

func ListSort(head *ListNode, tail *ListNode) *ListNode {
    if head == nil {
        return head
    }

    if head.Next == tail {
        head.Next = nil
        return head
    }

    slow, fast := head, head

    for fast != tail {
        fast = fast.Next
        slow = slow.Next
        if fast != tail {
            fast = fast.Next
        }
    }
    mid := slow
    return mergeList(ListSort(head, mid), ListSort(mid, tail))
}

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    dummy := res

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            dummy.Next = l1
            dummy = dummy.Next
            l1 = l1.Next
        } else {
            dummy.Next = l2
            dummy = dummy.Next
            l2 = l2.Next
        }
    }

    if l1 != nil {
        dummy.Next = l1
    }
    if l2 != nil {
        dummy.Next = l2
    }
    return res.Next
}

func copyRandomList(head *Node) *Node {
    cur := head

    mymap := make(map[*Node]*Node)

    for cur != nil {
        mymap[cur] = &Node{Val: cur.Val}
        cur = cur.Next
    }
    cur = head

    for cur != nil {
        mymap[cur].Next = mymap[cur.Next]
        mymap[cur].Random = mymap[cur.Random]
        cur = cur.Next
    }

    return mymap[head]
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    // 1 2 3 4 5
    // 2 1 4 3 5
    dummy := &ListNode{Val: -1}
    dummy.Next = head

    pre := dummy
    end := dummy

    for end.Next != nil {
        for i := 0; i < k && end != nil; i++ {
            end = end.Next
        }
        if end == nil {
            break
        }
        start, next := pre.Next, end.Next
        end.Next = nil

        pre.Next = helper(start)

        start.Next = next
        pre = start
        end = pre
    }
    return dummy.Next
}

func helper(head *ListNode) *ListNode {

    var pre *ListNode

    for head != nil {
        next := head.Next
        head.Next = pre
        pre = head
        head = next
    }
    return pre
}

func swapPairs(head *ListNode) *ListNode {
    // 1 2 3 4
    // 2 1 4 3
    dummy := &ListNode{Val: -1}
    dummy.Next = head
    pre := dummy
    cur := head

    for cur != nil && cur.Next != nil {
        next := cur.Next.Next
        pre.Next = cur.Next
        cur.Next.Next = cur
        cur.Next = next

        pre = cur
        cur = next
    }

    return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 1 2 3 4 5 6
    // 0 1 2 3 4 5 6
    dummy := &ListNode{Val: 0}
    dummy.Next = head
    fast, slow := dummy, dummy

    for i := 0; i < n; i++ {
        fast = fast.Next
    }

    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    res := &ListNode{}
    dummy := res

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            dummy.Next = list1
            dummy = dummy.Next
            list1 = list1.Next
        } else {
            dummy.Next = list2
            dummy = dummy.Next
            list2 = list2.Next
        }
    }

    if list1 != nil {
        dummy.Next = list1
    }
    if list2 != nil {
        dummy.Next = list2
    }
    return res.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // l1 = [2,4,3], l2 = [5,6,4]
    // [7,0,8]
    // 342 + 465 = 807
    res := &ListNode{}
    dummy := res
    x, y, div, sum := 0, 0, 0, 0
    for l1 != nil || l2 != nil || div != 0 {
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        } else {
            x = 0
        }
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        } else {
            y = 0
        }

        sum = x + y + div
        div = sum / 10
        dummy.Next = &ListNode{Val: sum % 10}
        dummy = dummy.Next
    }

    return res.Next
}
