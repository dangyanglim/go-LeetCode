package main
import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
	     Val int
	     Next *ListNode
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func addList(head *ListNode,node *ListNode){
	var temp *ListNode=head
	for {
		if  temp.Next!=nil{
			temp=temp.Next
		}else{
			break;
		}
	}
	temp.Next=node
}
func showList(head *ListNode){
	var temp *ListNode=head
	fmt.Printf("showList\n")
	for {
		fmt.Println(temp.Val)
		if  temp.Next!=nil{
			temp=temp.Next
		}else{
			break;
		}
	}
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var sta *ListNode=head
    var sec *ListNode=head
    for i:=0;i<n;i++ {
        if sta.Next!=nil {
            sta=sta.Next
        }else {
            return head.Next
        }
    }
    for {
        if sta.Next!=nil {
            sta=sta.Next
            sec=sec.Next
        }else {
            if sec.Next!=nil{
               sec.Next=sec.Next.Next 
            }else{
                sec=nil
            }
            
            break
        }
    }
    return head
}
func reverseList(head *ListNode) *ListNode {
	var p *ListNode=head
	if head==nil {
		return head
	}
	var q *ListNode=head.Next
	var r *ListNode
	head.Next=nil
    for {
		if q!=nil{
			r=q.Next
			q.Next=p
			p=q
			q=r
		}else{
			break
		}
	}
	head=p
	return head
}
func reverseList2(head *ListNode) *ListNode {
	var newhead *ListNode
	if head==nil||head.Next==nil{
		return head
	}
	newhead=reverseList2(head.Next)
	head.Next.Next=head
	head.Next=nil
	return newhead
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head:=&ListNode{
		Val:5,
	}
	var newNode *ListNode=head
	for ;l1!=nil&&l2!=nil;{
		if l1.Val>l2.Val{
			newNode.Next=l2
			newNode=l2
			l2=l2.Next
		}else {
			newNode.Next=l1
			newNode=l1
			l1=l1.Next			
		}
	}
	if l1==nil{
		newNode.Next=l2
	}
	if l2==nil{
		newNode.Next=l1
	}
	return head.Next
}
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	head:=&ListNode{
		Val:5,
	}
	var newNode *ListNode=head
	if l1==nil {
		return l2
	}
	if l2==nil {
		return l1
	}
	if l1.Val<l2.Val {
		newNode=l1
		newNode.Next=mergeTwoLists2(l1.Next,l2)
	}else{
		newNode=l2
		newNode.Next=mergeTwoLists2(l2.Next,l1)		
	}
	return newNode
}
func isPalindrome(head *ListNode) bool {
    if head==nil||head.Next==nil{
		return true
	}
	var fast *ListNode=head.Next
	var slow *ListNode=head
	for ;fast!=nil&&fast.Next!=nil;{
		fast=fast.Next.Next
		slow=slow.Next
	}
	var secondHead *ListNode=nil
	var second *ListNode=slow.Next
	var temp *ListNode
	for ;second!=nil;{
		temp=second.Next
		second.Next=secondHead
		secondHead=second
		second=temp
	}
	for ;secondHead!=nil&&secondHead.Val==head.Val;{
		secondHead=secondHead.Next
		head=head.Next
	}
	if secondHead==nil{
		return true
	}
	return false
}
func hasCycle(head *ListNode) bool {
    if head==nil||head.Next==nil{
		return false
	}
	var fast *ListNode=head.Next
	var slow *ListNode=head
	for ;fast!=nil&&fast.Next!=nil;{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow {
			return true
		}
	}
	return false
}
func maxDepth(root *TreeNode) int {
    if root==nil {
		return 0
	}
	if root.Left==nil&&root.Right==nil {
		return 1
	}
	left:=maxDepth(root.Left)+1
	right:=maxDepth(root.Right)+1
	if left>right {
		return left
	}else{
		return right
	}

}

func isValidBST(root *TreeNode) bool {
	INT_Max := int(^uint64(0) >> 1)
    INT_Min := ^INT_Max
	if root==nil{
		return true
	}
	return isValid(root,INT_Min,INT_Max)
}
func isValid(root *TreeNode,low ,high int) bool{
	if root==nil{
		return true
	}
	if root.Val<=low||root.Val>=high{
		return false
	}
	return isValid(root.Left,low,root.Val)&&isValid(root.Right,root.Val,high)
}
func isSymmetric(root *TreeNode) bool {
    if root==nil{
        return true
    }
    return isSameTree(root.Left,invert(root.Right))
}
func invert(root *TreeNode) *TreeNode{
	if root==nil{
		return nil
	}
	node:= root.Left
	root.Left=invert(root.Right)
	root.Right=invert(node)
	return root
}
func isSameTree(q *TreeNode,p *TreeNode) bool{
	if q==nil||p==nil{
		return q==p
	}
    if q!=nil&&p!=nil{
		if q.Val!=p.Val{
			return false
		}
	}
    if isSameTree(q.Left,p.Left)&&isSameTree(p.Right,q.Right){
		return true
	}
	return false
}
func levelOrder(root *TreeNode) [][]int {
    if root==nil{
		return [][]int{}
	}
	res:=[][]int{}
	queue:=[]*TreeNode{root}
	for ;len(queue)>0;{
		levelSize:=len(queue)
		currentLevel:=[]int{}
		for i:=0;i<levelSize;i++{
			node:=queue[0]
			queue=queue[1:]
			currentLevel=append(currentLevel,node.Val)
			if node.Left!=nil{
				queue=append(queue,node.Left)
			}
			if node.Right!=nil{
				queue=append(queue,node.Right)
			}
		}
		res=append(res,currentLevel)
	}
	return res
}
func sortedArrayToBST(nums []int) *TreeNode {
    return sortedToBST(nums,0,len(nums)-1)
}
func sortedToBST(nums []int,left int, right int) *TreeNode {
	if left>right {
		return nil
	}
	mid:=(left+right)/2
	cur:=&TreeNode{
		Val:nums[mid],
	}
	cur.Left=sortedToBST(nums,left,mid-1)
	cur.Right=sortedToBST(nums,mid+1,right)
	return cur
    
}
func main() {
	head:=&ListNode{
		Val:5,
	}
	tree:=&TreeNode{
		Val:2147483647,
	}
	// head5:=&ListNode{
	// 	Val:1,
	// }
	node1:=&ListNode{
		Val:2,
	}
	// node2:=&ListNode{
	// 	Val:4,
	// }
	// node3:=&ListNode{
	// 	Val:2,
	// }
	// node4:=&ListNode{
	// 	Val:4,
	// }
	 addList(head,node1)
	 //addList(head,node2)
	//addList(head5,node3)
	//addList(head5,node4)
	showList(head)
	//showList(head5)
	//removeNthFromEnd(head,1)
	//head=reverseList(head)
	//head=reverseList2(head)
	//head=mergeTwoLists2(head5,head)
	//fmt.Println(hasCycle(head))
	fmt.Println(levelOrder(tree))
	fmt.Println(isValidBST(tree))
	showList(head)
}
