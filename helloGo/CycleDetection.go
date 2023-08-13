package main

import "fmt"

func main() {
	vals := []int{1, 2, 3, 4, 5, 7, 8, 9}
	head := buildCircularLinkedList(vals)
	res := detectCycle(head)
	fmt.Println(res.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil // 无环
		}
		fast = fast.Next.Next
		if fast == slow { // 有环、初次相遇
			// a + (n+1)b + nc = 2(a+b) ⟹ a=c + (n−1)(b+c)
			// 因此，当发现 slow 与 fast 相遇时，我们再额外使用一个指针 ptr。
			//  起始，它指向链表头部；随后，它和 slow 每次向后移动一个位置。最终，它们会在入环点相遇。
			p := head // 另一只🐢从头开始
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// 构建环形链表
func buildCircularLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	// 创建头节点
	head := &ListNode{Val: vals[0]}
	current := head
	cross := head // ❌节点

	// 创建链表节点并连接形成环
	for i := 1; i < len(vals); i++ {
		node := &ListNode{Val: vals[i]}
		current.Next = node
		current = node
		if i == 3 {
			cross = node
		}
	}

	// 将最后一个节点的 Next 指向cross节点，形成环
	current.Next = cross
	return head
}
