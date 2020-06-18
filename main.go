package main

import (
	"fmt"
	ln "github.com/soeyi/leetcode-go/21.MergeTwoSortedLists"

	pl "github.com/soeyi/leetcode-go/234.PalindromeLinkedList"
)

func main() {
	l1 := &ln.ListNode{Val: 1, Next: &ln.ListNode{Val: 2, Next: &ln.ListNode{Val: 3}}}
	l2 := &ln.ListNode{Val: 2, Next: &ln.ListNode{Val: 4, Next: &ln.ListNode{Val: 5}}}

	res := *ln.MergeTwoLists(l1, l2)

	fmt.Printf("res &+v", res)

	l := &pl.ListNode{Val: 1, Next: &pl.ListNode{Val: 2, Next: &pl.ListNode{Val: 1}}}

	c := pl.IsPalindrome(l)

	fmt.Println("")

	fmt.Println("c", c)
}
