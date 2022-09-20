package _00HotTopic

import (
	"strconv"
	"strings"
)

// 297. 二叉树的序列化与反序列化
// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/?favorite=2cktkvj

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 初始化字符串
	result := []byte{}

	// 对二叉树进行先序遍历
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		// 遇到空节点
		if root == nil {
			result = append(result, "null,"...)
			return
		}
		// 当前节点不为空，将 val 追加到result
		result = append(result, strconv.Itoa(root.Val)...)
		result = append(result, ',')
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)
	// 去除末尾的逗号
	result = result[:len(result)-1]
	return string(result)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	source := strings.Split(data, ",")
	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		// 节点值为空
		if source[0] == "null" {
			source = source[1:]
			return nil
		}
		// 节点值不为空
		val, _ := strconv.Atoi(source[0])
		source = source[1:]
		return &TreeNode{val, buildTree(), buildTree()}
	}
	return buildTree()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
