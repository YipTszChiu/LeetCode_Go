package SearchAndBackchecking

import (
	"strconv"
	"strings"
)

//剑指 Offer 37. 序列化二叉树
//https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof/

func serialize(root *TreeNode) string {
	ans := "["

	//中序遍历
	que := []*TreeNode{root}

	for len(que) > 0 {
		//出队
		temp := que[0]
		que = que[1:]
		//处理成字符串
		if temp == nil {
			ans += "nill,"
			continue
		} else {
			ans += strconv.Itoa(temp.Val) + ","
		}
		//将左右节点入队
		que = append(que, temp.Left)
		que = append(que, temp.Right)
	}
	//将末尾的","去掉
	ans = ans[:len(ans)-1]

	return ans + "]"
}

func deserialize(data string) *TreeNode {
	//将左右的"[" 和 "]"去除
	data = data[1 : len(data)-1]
	//使用Split函数分割字符串
	dataSplit := strings.Split(data, ",")
	//新建一个空队列
	que := []*TreeNode{}
	//创建根节点
	dataSplitCount := 0
	rootVal, _ := strconv.Atoi(dataSplit[dataSplitCount])
	root := &TreeNode{rootVal, nil, nil}
	dataSplitCount++
	//根节点入队
	que = append(que, root)

	for len(que) > 0 {
		//出队
		temp := que[0]
		que = que[1:]

		//添加左节点
		//获取数组中对应层次序列的内容
		leftNodeContext := dataSplit[dataSplitCount]
		dataSplitCount++
		//如果为空
		if leftNodeContext == "nil" {
			temp.Left = nil
		} else {
			//如果不为空，创建一个左节点
			leftVal, _ := strconv.Atoi(leftNodeContext)
			leftNode := &TreeNode{leftVal, nil, nil}
			//当前节点的左指针指向左节点
			temp.Left = leftNode
			//左节点入队
			que = append(que, leftNode)
		}

		//添加右节点
		//获取数组中对应层次序列的内容
		rightNodeContext := dataSplit[dataSplitCount]
		dataSplitCount++
		//如果为空
		if rightNodeContext == "nil" {
			temp.Right = nil
		} else {
			//如果不为空，创建一个右节点
			rightVal, _ := strconv.Atoi(rightNodeContext)
			rightNode := &TreeNode{rightVal, nil, nil}
			//当前节点的右指针指向右节点
			temp.Right = rightNode
			//右节点入队
			que = append(que, rightNode)
		}
	}

	return root
}
