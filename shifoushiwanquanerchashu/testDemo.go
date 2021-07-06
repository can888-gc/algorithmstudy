package main

import "fmt"

/**
 *@DemoName: 判断是否是完全二叉树
 *@Description: 判断是否是完全二叉树
 *@Author: can888-gc
 *@Date: 2021/7/6 下午10:26
 */

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func main() {
	root := &TreeNode{val: "1"}
	root.left = &TreeNode{val: "2"}
	root.left.left = &TreeNode{val: "4"}
	//root.left.left.left = &TreeNode{val: "7"}
	root.right = &TreeNode{val: "3"}
	//root.right.left = &TreeNode{val: "5"}
	//root.right.right = &TreeNode{val: "6"}
	if IsCompleteBt(root) {
		fmt.Println("是完全二叉树")
	} else {
		fmt.Println("不是完全二叉树")
	}
}

// IsCompleteBt 这里默认根节点为空属于完全二叉树,这个可以自已定义是否为完全二叉树/***/
func IsCompleteBt(root *TreeNode) bool {
	if root == nil {
		return true
	}

	/**
	* 条件:
	* 	1.当一个节点存在右子节点但是不存在左子节点这颗树视为非完全二叉树
	*	2.当一个节点的左子节点存在但是右子节点不存在视为完全二叉树
	 */
	var tempNodeQueue []*TreeNode

	tempNodeQueue = append(tempNodeQueue, root)

	var tempNode *TreeNode

	isSingleNode := false
	for len(tempNodeQueue) != 0 {
		tempNode = tempNodeQueue[0]
		tempNodeQueue = tempNodeQueue[1:]

		if tempNode.left == nil {
			isSingleNode = true
		} else {
			tempNodeQueue = append(tempNodeQueue, tempNode.left)
		}

		if tempNode.right == nil {
			isSingleNode = true
		} else {
			tempNodeQueue = append(tempNodeQueue, tempNode.right)
		}

		if (isSingleNode && (tempNode.right != nil || tempNode.left != nil)) || (tempNode.right != nil && tempNode.left == nil) {
			return false
		}
	}
	return true
}
