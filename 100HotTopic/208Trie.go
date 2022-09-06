package _00HotTopic

// 208.实现前缀树
// https://leetcode.cn/problems/implement-trie-prefix-tree/solution/shi-xian-trie-qian-zhui-shu-by-leetcode-ti500/

type Trie struct {
	// 英文字符26个，需要一个26位的指针
	children [26]*Trie
	// 表示是否为一个单词的末尾
	isEnd bool
}

// 构造器：返回一个空的前缀树
func Constructor_Trie() Trie {
	return Trie{}
}

// Insert 实现插入一个新的单词
func (t *Trie) Insert(word string) {
	node := t
	// 遍历单词
	for _, c := range word {
		// 指针为 0~25，需要转换
		c -= 'a'
		// 如果对应字符位为空
		if node.children[c] == nil {
			// 创建一颗新的Trie
			node.children[c] = &Trie{}
		}
		// node 指向对应字符构成的 Trie
		node = node.children[c]
	}
	// 遍历结束后将 isEnd 位置为 true
	node.isEnd = true
}

// SearchPrefix 查找是否有某个前缀
func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, c := range prefix {
		// 指针为 0~25，需要转换
		c -= 'a'
		if node.children[c] == nil {
			return nil
		}
		node = node.children[c]
	}
	return node
}

// Search 查找 Trie 中是否含 word 这个单词
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith 调用 SearchPrefix 查找是否有某个前缀
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
