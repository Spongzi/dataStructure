package main

import (
	"fmt"
	"math"
)

// Node 定义节点结构体
type Node struct {
	value  int   // 值
	left   *Node // 左节点
	right  *Node // 右节点
	parent *Node // 父节点
}

// Tree 定义树结构体
type Tree struct {
	root   *Node // 跟
	length int   // 长度
}

// 调用创建树的方法
func main() {
	creatTree()
}

func creatTree() {
	// 定义一个列表
	arrList := []int{14, 2, 5, 7, 23, 35, 12, 17, 31}
	// 实现一个树的实体
	myTree := Tree{}
	for i := 0; i < len(arrList); i++ {
		// 使用插入节点
		myTree = insertNode(myTree, arrList[i])
		myTree.length++
	}
	fmt.Println(myTree)
	//LDR(myTree)
	TreeHeight(myTree)
}

// TreeHeight 树的高度
func TreeHeight(tree Tree) {
	// 定义一个左边节点的高度
	var hl = 1
	// 如果左边节点的值不为空
	if tree.root.left != nil {
		// 左边的高度调用高度的方法
		hl = heightMax(tree.root.left, hl)
	}
	// 定义右边节点的高度
	var hr = 1
	// 如果右边的节点不为空
	if tree.root.right != nil {
		hr = heightMax(tree.root.left, hr)
	}
	fmt.Println(hl, hr)
	// 找出最高值
	fmt.Println("Tree height is ", int(math.Max(float64(hl), float64(hr))))
}

func heightMax(node *Node, h int) int {
	var hL = h
	var hR = h
	// 如果说左边或右边的节点都为空，那么就答应node节点，直接返回h
	if node.left == nil && node.right == nil {
		fmt.Println(node)
		return h
	}
	// 如果左边的节点不为空
	if node.left != nil {
		// 高度加一
		h++
		// 继续调用heightMax直到左右节点都为空
		hL = heightMax(node.left, h)
	}
	// 如果右节点不为空
	if node.right != nil {
		h++
		// 继续调用heightMax直到左右节点都为空
		hR = heightMax(node.right, h)
	}
	return int(math.Max(float64(hL), float64(hR)))
}

// LDR 中序遍历
func LDR(tree Tree) {
	readList := make(map[int]int)
	i := 0
	var currentNode *Node
	currentNode = tree.root
	for {
		//fmt.Println(currentNode)
		if i == tree.length {
			//fmt.Println(currentNode.value)
			break
		}
		if currentNode.left == nil {
			if readList[currentNode.value] == 1 {
				if readList[currentNode.right.value] == 1 {
					currentNode = currentNode.parent
					continue
				} else {
					currentNode = currentNode.right
					continue
				}
			} else {
				fmt.Println(currentNode.value)
				readList[currentNode.value] = 1
				i++
				if currentNode.right == nil {
					currentNode = currentNode.parent
					continue
				} else {
					if readList[currentNode.right.value] == 1 {
						currentNode = currentNode.parent
						continue
					} else {
						currentNode = currentNode.right
						continue
					}
				}
			}
		} else {
			if readList[currentNode.left.value] == 1 {
				if readList[currentNode.value] == 1 {
					currentNode = currentNode.right
					continue
				} else {
					fmt.Println(currentNode.value)
					readList[currentNode.value] = 1
					i++
					if currentNode.right == nil {
						currentNode = currentNode.parent
						continue
					} else {
						if readList[currentNode.right.value] == 1 {
							currentNode = currentNode.parent
							continue
						} else {
							currentNode = currentNode.right
							continue
						}

					}
				}
			} else {
				currentNode = currentNode.left
				continue
			}

		}

	}
}

// 插入节点
func insertNode(tree Tree, insertValue int) Tree {
	var currentNode *Node
	var tmp *Node
	i := 0
	if tree.length == 0 {
		currentNode = new(Node)
		currentNode.value = insertValue
		tree.root = currentNode
		return tree
	} else {
		currentNode = tree.root
	}
	for {
		//fmt.Println(currentNode)
		if currentNode.value < insertValue {
			//判断是否有右节点
			if currentNode.right == nil {
				tmp = new(Node)
				tmp.value = insertValue
				currentNode.right = tmp
				tmp.parent = currentNode
				break
			} else {
				currentNode = currentNode.right
				continue
			}
		} else {
			if currentNode.left == nil {
				tmp = new(Node)
				tmp.value = insertValue
				currentNode.left = tmp
				tmp.parent = currentNode
				break
			} else {
				currentNode = currentNode.left
				continue
			}
		}
		i++
	}
	return tree
}
