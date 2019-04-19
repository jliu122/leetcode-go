package list

// 002
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{Val:-1}
	curr := dummy
	carry := 0
	for l1 != nil && l2 != nil {
		carry += l1.Val + l2.Val
		curr.Next = &ListNode{Val:carry % 10}
		curr = curr.Next
		carry /= 10
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		l1 = l2
	}
	for l1 != nil {
		carry += l1.Val
		curr.Next = &ListNode{Val:carry % 10}
		curr = curr.Next
		carry /= 10
		l1 = l1.Next
	}
	if carry == 1 {
		curr.Next = &ListNode{Val:1}
	}
	return dummy.Next
}
