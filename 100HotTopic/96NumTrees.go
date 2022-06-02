package _00HotTopic

// 96. 不同的二叉搜索树
// https://leetcode.cn/problems/unique-binary-search-trees/

func numTrees(n int) int {
	// 创建一个 dp 数组用于存储结果
	G := make([]int, n+1)
	// G(n) 表示：n 个节点构成不同的二叉搜索树数量
	// F(i, n) 表示：总共 n 个节点，根节点为 i 构成不同二叉搜索树的数量
	// * F(i, n) = G(i-1) * G(n-i)   i为根，左子树 i-1 个元素构成不同二叉树，右子树同理
	// ** G(n) 则需要统计 F(1) + F(2) + …… + F(n)，表示n个节点构成不同二叉树，1~n都可以作为根
	// 初始值
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j] // 等号右边是以 j 为根的二叉树数量
		}
	}

	return G[n]
}
