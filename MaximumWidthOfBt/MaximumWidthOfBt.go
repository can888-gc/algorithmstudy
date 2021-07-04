package main

import "fmt"

/**
 *@DemoName: 查找二叉树的最大宽度
 *@Description: 查找二叉树的最大宽度
 *@Author: can888-gc
 *@Date: 2021/7/4 上午11:59
 */

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func main() {
	sNode := &TreeNode{val: "1"}
	sNode.left = &TreeNode{val: "2"}
	sNode.right = &TreeNode{val: "3"}
	sNode.left.left = &TreeNode{val: "4"}
	sNode.left.right = &TreeNode{val: "5"}
	sNode.right.left = &TreeNode{val: "6"}
	sNode.left.left.left = &TreeNode{val: "7"}
	sNode.left.left.right = &TreeNode{val: "8"}
	sNode.left.left.left.left = &TreeNode{val: "9"}
	sNode.left.left.left.right = &TreeNode{val: "10"}
	sNode.left.left.right.left = &TreeNode{val: "11"}
	sNode.left.left.right.right = &TreeNode{val: "12"}

	sNode.left.right.left = &TreeNode{val: "13"}
	sNode.left.right.right = &TreeNode{val: "14"}
	sNode.right.left.left = &TreeNode{val: "15"}
	sNode.right.left.right = &TreeNode{val: "16"}
	maxW, row := findBtMaxWidth(sNode)
	fmt.Printf("最大宽度: %v;在第 %v层", maxW, row)
}

func findBtMaxWidth(bt *TreeNode) (maxWidth int, maxInrow int) {
	row := 0
	//临时保存节点的队列
	var tempSaveNodeQueue []*TreeNode
	//保存宽度
	count := 1
	var currentRowEndNode *TreeNode
	var nextRowEndNode *TreeNode
	if bt != nil {
		nextRowEndNode = bt
		currentRowEndNode = nextRowEndNode
		tempSaveNodeQueue = append(tempSaveNodeQueue, bt)
	}

	for len(tempSaveNodeQueue) != 0 {
		count++
		treeNode := tempSaveNodeQueue[0]
		tempSaveNodeQueue = tempSaveNodeQueue[1:]

		if treeNode.left != nil {
			nextRowEndNode = treeNode.left
			tempSaveNodeQueue = append(tempSaveNodeQueue, treeNode.left)
		}

		if treeNode.right != nil {
			nextRowEndNode = treeNode.right
			tempSaveNodeQueue = append(tempSaveNodeQueue, treeNode.right)
		}

		if currentRowEndNode == treeNode {
			row++
			currentRowEndNode = nextRowEndNode
			if maxWidth < count {
				maxInrow = row
				maxWidth = count
			}
			count = 0
		}
	}
	return
}
