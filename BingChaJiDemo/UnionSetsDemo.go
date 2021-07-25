package main

type Node struct {
	Val string
}

type UnionSet struct {
	Nodes   map[string]Node
	Parents map[Node]Node
	SizeMap map[Node]int
}

/**
初始化map数据
*/
func (u *UnionSet) InitUnionSet(values []string) {
	u.Nodes = make(map[string]Node)
	u.Parents = make(map[Node]Node)
	u.SizeMap = make(map[Node]int)

	for _, value := range values {
		n := &Node{Val: value}
		u.Nodes[value] = *n
		u.Parents[*n] = *n
		u.SizeMap[*n] = 1
	}
}

func (u *UnionSet) unionNode(nodePre string, nodeSuffix string) {
	tempAncestorNodePre := u.findAncestorNode(nodePre)
	tempAncestorNodeSuffix := u.findAncestorNode(nodeSuffix)

	if tempAncestorNodePre != tempAncestorNodeSuffix {
		nodePreSize := u.SizeMap[tempAncestorNodePre]
		nodeSuffixSize := u.SizeMap[tempAncestorNodeSuffix]

		if nodePreSize >= nodeSuffixSize {
			u.Parents[tempAncestorNodeSuffix] = tempAncestorNodePre
			u.SizeMap[tempAncestorNodePre] = nodePreSize + nodeSuffixSize
			delete(u.SizeMap, tempAncestorNodeSuffix)
		} else {
			u.Parents[tempAncestorNodePre] = tempAncestorNodeSuffix
			u.SizeMap[tempAncestorNodeSuffix] = nodePreSize + nodeSuffixSize
			delete(u.SizeMap, tempAncestorNodePre)
		}
	}
}

/**
查找祖先节点
*/
func (u *UnionSet) findAncestorNode(n string) Node {
	tempNode := u.Nodes[n]
	for tempNode != u.Parents[tempNode] {
		tempNode = u.Parents[tempNode]
	}
	return tempNode
}

func main() {
	values := []string{"2", "3", "4", "5"}
	u := &UnionSet{}
	u.InitUnionSet(values)
	u.unionNode("3", "4")
}
