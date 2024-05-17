package main

import (
	"fmt"
)

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

// CheckDuplicateIDs returns the value and level of the shallowest duplicate node.
func CheckDuplicateIDs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	type nodeInfo struct {
		node  *TreeNode
		level int
	}

	queue := []nodeInfo{{node: root, level: 1}}
	seen := make(map[int]int)
	shallowestDuplicate := 0
	minLevel := int(^uint(0) >> 1) // Max int value

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if _, exists := seen[current.node.Value]; exists {
			if current.level < minLevel {
				minLevel = current.level
				shallowestDuplicate = current.node.Value
			}
		} else {
			seen[current.node.Value] = current.level
		}

		for _, child := range current.node.Children {
			queue = append(queue, nodeInfo{node: child, level: current.level + 1})
		}
	}

	if minLevel == int(^uint(0)>>1) {
		return 0, 0
	}
	return shallowestDuplicate, minLevel
}

func main() {
	// Example usage:
	root := &TreeNode{
		Value: 1,
		Children: []*TreeNode{
			{Value: 2, Children: []*TreeNode{
				{Value: 4},
				{Value: 5},
			}},
			{Value: 3, Children: []*TreeNode{
				{Value: 6},
				{Value: 2}, // Duplicate
			}},
		},
	}

	value, level := CheckDuplicateIDs(root)
	fmt.Printf("Shallowest Duplicate: %d at Level: %d\n", value, level)
}
