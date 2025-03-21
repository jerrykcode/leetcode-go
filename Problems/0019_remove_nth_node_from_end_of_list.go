/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var nodes = []*ListNode{}
    for p := head; p != nil; p = p.Next {
        nodes = append(nodes, p)
    }
    tot := len(nodes)
    del_idx := tot - n
    if del_idx == 0 {
        head = nodes[0].Next
    } else {
        nodes[del_idx - 1].Next = nodes[del_idx].Next
    }
    nodes[del_idx].Next= nil
    return head
}