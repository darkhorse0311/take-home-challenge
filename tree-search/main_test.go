package main

import (
	"testing"
)

func TestCheckDuplicateIDs(t *testing.T) {
	tests := []struct {
		no            int
		name          string
		root          *TreeNode
		expectedValue int
		expectedLevel int
	}{
		{
			no:   1,
			name: "Test with duplicates",
			root: &TreeNode{
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
			},
			expectedValue: 2,
			expectedLevel: 3,
		},
		{
			no:   2,
			name: "Test without duplicates",
			root: &TreeNode{
				Value: 1,
				Children: []*TreeNode{
					{Value: 2, Children: []*TreeNode{
						{Value: 4},
						{Value: 5},
						{
							Value: 8,
							Children: []*TreeNode{
								{Value: 9},
								{Value: 10},
							},
						},
					}},
					{Value: 3, Children: []*TreeNode{
						{Value: 6},
						{Value: 7},
					}},
				},
			},
			expectedValue: 2,
			expectedLevel: 3,
		},
		{
			no:   3,
			name: "Test without duplicates",
			root: &TreeNode{
				Value: 1,
				Children: []*TreeNode{
					{Value: 2},
					{Value: 3},
				},
			},
			expectedValue: 0,
			expectedLevel: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, level := CheckDuplicateIDs(tt.root)
			if value != tt.expectedValue || level != tt.expectedLevel {
				t.Errorf("No: %d, Result: %d, %d; Expect: %d, %d", tt.no, value, level, tt.expectedValue, tt.expectedLevel)
			} else {
				t.Logf("No: %d, Result: %d, %d; Expect: %d, %d", tt.no, value, level, tt.expectedValue, tt.expectedLevel)
			}
		})
	}
}
