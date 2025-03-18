/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var p1, p2, res, tail *ListNode = l1, l2, nil, nil
    var k int = 0
    for p1 != nil || p2 != nil {
        val := k
        if p1 != nil {
            val += p1.Val
            p1 = p1.Next
        }
        if p2 != nil {
            val += p2.Val
            p2 = p2.Next
        }
        k = val / 10
        val %= 10
        tail = insert(tail, val)
        if res == nil {
            res = tail
        }
    }
    if k != 0 {
        tail = insert(tail, k)
        if res == nil {
            res = tail
        }
    }
    return res
}

func insert(tail *ListNode, val int) *ListNode {
    var p *ListNode = &ListNode{Val:val, Next:nil}
    if tail != nil {
        tail.Next = p
    } else {
        tail = p
    }
    return p
}