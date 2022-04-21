package SearchAndBackchecking

//剑指 Offer 12. 矩阵中的路径
//https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/

func exist(board [][]byte, word string) bool {
	//深度优先遍历
	row := len(board)
	column := len(board[0])

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false

}

func dfs(board [][]byte, word string, i, j, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	//word已遍历完成，匹配成功
	if k == len(word)-1 {
		return true
	}
	//置零当前位，避免在同一个位置反复遍历
	board[i][j] = 0
	//dfs查找当前位置的左右上下
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j+1, k+1) || dfs(board, word, i, j-1, k+1)
	//复原当前位
	board[i][j] = word[k]
	return res
}
