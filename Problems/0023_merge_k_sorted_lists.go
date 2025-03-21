/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    var res, res_tail *ListNode = nil, nil
    var min_list **ListNode = nil
    lists_num := len(lists)
    for lists_num > 0 {
        min_list = nil
        for i := 0; i < len(lists); i++ {
            if lists[i] != nil {
                if min_list == nil || lists[i].Val < (*min_list).Val {
                    min_list = &lists[i]
                }
            }
        }
        if min_list == nil {
            break
        }
        if res == nil {
            res = *min_list
            res_tail = *min_list
        } else {
            res_tail.Next = *min_list
            res_tail = *min_list
        }
        *min_list = (*min_list).Next
        if *min_list == nil {
            lists_num--
        }
    }
    return res
}