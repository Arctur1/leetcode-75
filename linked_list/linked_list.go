package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) traverse_linked_list() []int {
	list := l
	result := []int{}
	for {
		result = append(result, list.Val)
		if list.Next == nil {
			break
		}
		list = list.Next
	}
	return result
}

func (l *ListNode) node(n int) *ListNode {
	node := l
	for i := 0; i <= n; i++ {
		node = l.Next
	}
	return node
}

func create_linked_list(arr []int) *ListNode {
	head := new(ListNode)
	node := head
	for i := 0; i < len(arr); i++ {
		node.Val = arr[i]
		if i < len(arr)-1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}

	return head
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				fast, slow = fast.Next, slow.Next
			}
			return slow
		}
	}

	return nil

}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	reversed := &ListNode{head.Val, nil}

	for {
		head = head.Next
		reversed = &ListNode{head.Val, reversed}
		if head.Next == nil {
			break
		}
	}

	return reversed
}

func linkedListLength(head *ListNode) int {
	if head == nil {
		return 0
	}

	length := 1
	for head.Next != nil {
		head = head.Next
		length += 1
	}
	return length
}

func getNodeByIndex(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head
}

func middleNode(head *ListNode) *ListNode {
	length := linkedListLength(head)

	return getNodeByIndex(head, length/2)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head

	for {

		if list1 == nil {
			curr.Next = list2
			break
		}

		if list2 == nil {
			curr.Next = list1
			break
		}

		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next

	}

	return head.Next
}

func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursive(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsRecursive(l1, l2.Next)
	return l2
}
