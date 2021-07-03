package main

import "fmt"

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func main() {
	//"1","2","3","null","null","4","5"
	dataQueue := []string{"1", "2", "3", "null", "null", "4", "5"}
	node := DeserializationTb(dataQueue)
	fmt.Println(node)

	sNode := &TreeNode{val: "1"}
	sNode.left = &TreeNode{val: "2"}
	sNode.right = &TreeNode{val: "3"}
	sNode.right.left = &TreeNode{val: "4"}
	sNode.right.right = &TreeNode{val: "5"}
	SerializationTb(sNode)
}

/**
序列化二叉树
*/
func SerializationTb(bt *TreeNode) {
	root := bt
	var tempQueue []*TreeNode
	var saveSerData []string
	if root != nil {
		tempQueue = append(tempQueue, root)
	}
	var tempNode *TreeNode
	for len(tempQueue) != 0 {
		tempNode = tempQueue[0]
		if tempNode != nil {
			saveSerData = append(saveSerData, tempNode.val)
		} else {
			saveSerData = append(saveSerData, "null")
		}
		tempQueue = tempQueue[1:]
		if tempNode != nil {
			tempQueue = append(tempQueue, tempNode.left)
			tempQueue = append(tempQueue, tempNode.right)
		}
	}
}

/**
反序列化二叉树
*/
func DeserializationTb(dataQueue []string) (resultNode *TreeNode) {
	if len(dataQueue) == 0 {
		return nil
	}
	var tempNodeQueue []*TreeNode
	resultNode = generateNode(dataQueue[0])
	dataQueue = dataQueue[1:]
	if resultNode != nil {
		tempNodeQueue = append(tempNodeQueue, resultNode)
	}
	var tempNode *TreeNode
	for len(tempNodeQueue) != 0 {
		tempNode = tempNodeQueue[0]
		tempNodeQueue = tempNodeQueue[1:]
		if len(dataQueue) > 0 {
			tempNode.left = generateNode(dataQueue[0])
			dataQueue = dataQueue[1:]
			tempNode.right = generateNode(dataQueue[0])
			dataQueue = dataQueue[1:]
		}
		if tempNode.left != nil {
			tempNodeQueue = append(tempNodeQueue, tempNode.left)
		}
		if tempNode.right != nil {
			tempNodeQueue = append(tempNodeQueue, tempNode.right)
		}
	}
	return
}

func generateNode(str string) *TreeNode {
	if str == "null" {
		return nil
	}
	return &TreeNode{val: str}
}
