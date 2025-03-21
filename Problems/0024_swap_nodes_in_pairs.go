/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var tail, p, n, nn *ListNode = head, head, head.Next, head.Next.Next
    for ; ; {
        if n != nil {
            n.Next = p
            if tail == head {
                head = n
            } else {
                tail.Next = n
            }
            tail = p
        } else {
            tail.Next = p
            tail = p
            break
        }
        p = nn
        if p == nil {
            break
        }
        n = p.Next
        if n == nil {
            nn = nil
        } else {
            nn = n.Next
        }
    }
    tail.Next = nil
    return head
}