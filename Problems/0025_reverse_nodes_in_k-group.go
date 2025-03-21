/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    len := 0
    for p := head; p != nil; p = p.Next {
        len++
    }
    if len < k || k == 1 {
        return head
    }
    var tail, p, n, nn *ListNode = head, head, head.Next, head.Next.Next
    cnt := 0
    for ; ; {
        if cnt + k <= len {
            tmp := p
            for i := 0; i < k - 1; i++ {
                n.Next = p
                p = n
                n = nn
                if n != nil {
                    nn = n.Next
                }
            }
            cnt += k
            if tail == head {
                head = p
            } else {
                tail.Next = p
            }
            tail = tmp
            p = n
            if p != nil {
                n = nn
                if n != nil {
                    nn = n.Next
                }
            }
        } else {
            tail.Next = p
            break
        }
    }
    return head
}