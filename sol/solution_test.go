package sol

import (
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree, cur *TreeNode
	hashMap := make(map[int]*TreeNode)
	for idx, val := range *input {
		num := 0
		if val != "null" {
			num, _ = strconv.Atoi(val)
		}
		if idx == 0 {
			tree = &TreeNode{Val: num}
			hashMap[num] = tree
		}
		if node, ok := hashMap[num]; ok {
			cur = node
		} else {
			cur = &TreeNode{Val: num}
			hashMap[num] = cur
		}
		if 2*idx+1 < len(*input) {
			// cur.Left
			if (*input)[2*idx+1] != "null" {
				left_val, _ := strconv.Atoi((*input)[2*idx+1])
				if left, exist := hashMap[left_val]; exist {
					cur.Left = left
				} else {
					left := &TreeNode{Val: left_val}
					hashMap[left_val] = left
					cur.Left = left
				}
			}
		}
		if 2*idx+2 < len(*input) {
			if (*input)[2*idx+2] != "null" {
				right_val, _ := strconv.Atoi((*input)[2*idx+2])
				if right, exist := hashMap[right_val]; exist {
					cur.Right = right
				} else {
					right := &TreeNode{Val: right_val}
					hashMap[right_val] = right
					cur.Right = right
				}
			}
		}
	}
	return tree
}

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "p = [1,2,3], q = [1,2,3]",
			args: args{p: CreateBinaryTree(&[]string{"1", "2", "3"}),
				q: CreateBinaryTree(&[]string{"1", "2", "3"})},
			want: true,
		},
		{
			name: "p = [1,2], q = [1,null,2]",
			args: args{p: CreateBinaryTree(&[]string{"1", "2"}),
				q: CreateBinaryTree(&[]string{"1", "null", "2"})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
