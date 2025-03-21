/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    p1, p2 := list1, list2
    if p1 == nil {
        return p2
    } else if p2 == nil {
        return p1
    }
    var p, tail *ListNode = nil, nil
    if p1.Val < p2.Val {
        p = p1
        p1 = p1.Next
    } else {
        p = p2
        p2 = p2.Next
    }
    tail = p
    for p1 != nil && p2 != nil {
        if p1.Val < p2.Val {
            tail.Next = p1
            p1 = p1.Next
        } else {
            tail.Next = p2
            p2 = p2.Next
        }
        tail = tail.Next
    }
    if p1 != nil {
        tail.Next = p1
    } else {
        tail.Next = p2
    }
    return p
}