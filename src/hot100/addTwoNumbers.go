package hot100

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	cur := head
	carry := 0
	for  l1 != nil || l2 != nil||carry>0 {
		sum:=carry
		if l1!=nil {
			sum+=l1.Val
			l1 = l1.Next
		}
		if l2!=nil {
			sum+=l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{sum % 10,nil}
		carry = sum / 10
		cur = cur.Next
	}
	return head.Next
}
