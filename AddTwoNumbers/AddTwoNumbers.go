/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry uint8
    head := NewListNode(0)
    tail := head

    for l1 != nil || l2 != nil || carry != 0 {
        var num1, num2 uint8

        if l1 != nil {
            num1 = uint8(l1.Val)
            l1 = l1.Next
        }

        if l2 != nil {
            num2 = uint8(l2.Val)
            l2 = l2.Next
        }

        sum := num1 + num2 + carry
        newVal := sum % 10
        carry = sum / 10

        newListNode := NewListNode(int(newVal))
        tail.Next = newListNode
        tail = tail.Next
    }
    return head.Next
}

func NewListNode(val int) *ListNode {
    return &ListNode{Val: val, Next: nil}
}