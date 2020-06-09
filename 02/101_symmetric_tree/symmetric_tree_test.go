package _101_symmetric_tree

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"tree0",
			args{&TreeNode{Val: 0, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}},
			true,
		},
		{
			"tree1",
			args{&TreeNode{Val: 0, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}},
			false,
		},
		{
			"tree2",
			args{nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
