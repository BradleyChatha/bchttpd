package main

import (
	"fmt"
)

type RadixNode struct {
	segment        [128]byte
	segmentLength  uint
	children       [64]uint
	childrenLength uint
	data           []byte
}

type RadixTree struct {
	nodes []RadixNode
}

func radixTreeCreate(initSize uint64) *RadixTree {
	tree := new(RadixTree)
	tree.nodes = make([]RadixNode, initSize)

	return tree
}

func (tree *RadixTree) insert(str string, data []byte) {
	if len(tree.nodes) == 0 {
		tree.nodes = append(tree.nodes, RadixNode{segment: [128]byte{}, segmentLength: 0, children: [64]uint{}, childrenLength: 0, data: []uint8{}})
	}

	currNode := &tree.nodes[0]
	strCursor := 0

While:
	for strCursor < len(str) {
		if str[strCursor] == currNode.segment[0] {
			strCursor++
		}
		for i := uint(0); i < currNode.childrenLength; i++ {
			child := &tree.nodes[currNode.children[i]]
			if str[strCursor] == child.segment[0] {
				currNode = child
				continue While
			}
		}
		break
	}

	if strCursor >= len(str) {
		return
	}

	for i := strCursor; i < len(str); i++ {
		newNode := RadixNode{childrenLength: 0, segmentLength: 1, segment: [128]byte{str[i]}, data: []byte{}}
		currNode.children[currNode.childrenLength] = uint(len(tree.nodes))
		currNode.childrenLength++

		tree.nodes = append(tree.nodes, newNode)
		currNode = &tree.nodes[len(tree.nodes)-1]
	}

	currNode.data = data
}

func (tree *RadixTree) consolidate(node *RadixNode, isRoot bool) {
	for i := uint(0); i < node.childrenLength; i++ {
		tree.consolidate(&tree.nodes[node.children[i]], false)
	}

	if node.childrenLength == 1 && len(node.data) == 0 && !isRoot {
		child := tree.nodes[node.children[0]]
		copy(node.segment[node.segmentLength:], child.segment[0:child.segmentLength])
		node.segmentLength += child.segmentLength
		node.data = child.data
		node.childrenLength = child.childrenLength
		node.children = child.children
	}
}

func (tree *RadixTree) finalise() {
	if len(tree.nodes) > 0 {
		tree.consolidate(&tree.nodes[0], true)
	}
}

func (tree *RadixTree) find(str string) []byte {
	currNode := &tree.nodes[0]
	strCursor := 0

While:
	for {
		charsLeft := len(str) - strCursor
		if charsLeft == 0 {
			return currNode.data
		} else if currNode.childrenLength == 0 {
			return []byte{}
		}

		for i := 0; i < len(currNode.children); i++ {
			child := &tree.nodes[currNode.children[i]]
			if child.segmentLength > 0 && child.segmentLength <= uint(charsLeft) && str[strCursor:uint(strCursor)+child.segmentLength] == string(child.segment[0:child.segmentLength]) {
				strCursor += int(child.segmentLength)
				currNode = child
				continue While
			}
		}

		return []byte{}
	}
}

func (tree *RadixTree) printNode(node *RadixNode, level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("-")
	}

	fmt.Printf("%s", string(node.segment[0:node.segmentLength]))
	if len(node.data) > 0 {
		fmt.Print("<")
	}
	fmt.Println()
	for i := uint(0); i < node.childrenLength; i++ {
		tree.printNode(&tree.nodes[node.children[i]], level+1)
	}
}

func (tree *RadixTree) print() {
	if len(tree.nodes) > 0 {
		tree.printNode(&tree.nodes[0], 0)
	}
}
