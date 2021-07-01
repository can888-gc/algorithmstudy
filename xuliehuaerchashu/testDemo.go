package main

import "fmt"

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func main() {
	dataQueue := []string{"5", "4", "null", "null", "3", "2", "1"}
	node := DeserializationTb(dataQueue)
	fmt.Println(node)
}

func DeserializationTb(dataQueue []string) (resultNode *TreeNode) {
	if len(dataQueue) == 0 {
		return nil
	}
	var tempNodeQueue []*TreeNode
	resultNode = generateNode(dataQueue[len(dataQueue) - 1])
	dataQueue = dataQueue[:len(dataQueue) - 1]
	if resultNode != nil {
		tempNodeQueue = append(tempNodeQueue,resultNode)
	}
	var tempNode *TreeNode
	for len(tempNodeQueue) != 0 {
		tempNode = tempNodeQueue[0]
		tempNodeQueue = tempNodeQueue[1:]
		if len(dataQueue) > 0 {
			tempNode.left = generateNode(dataQueue[len(dataQueue) - 1])
			dataQueue = dataQueue[:len(dataQueue) - 1]
			tempNode.right = generateNode(dataQueue[len(dataQueue) - 1])
			dataQueue = dataQueue[:len(dataQueue) - 1]
		}
		if tempNode.left != nil {
			tempNodeQueue = append(tempNodeQueue,tempNode.left)
		}
		if tempNode.right != nil {
			tempNodeQueue = append(tempNodeQueue,tempNode.right)
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
