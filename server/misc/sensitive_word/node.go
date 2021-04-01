package sensitive_word

const (
	// 每个字符map的初始化容量
	INIT_TRIE_CHILDREN_NUM = 128
)

type trieNode struct {
	// 是否为字符尾
	isEndOfWord bool
	// 后续字符map
	children map[rune]*trieNode
}

func newtrieNode() *trieNode {
	return &trieNode{
		isEndOfWord: false,
		children:    make(map[rune]*trieNode, INIT_TRIE_CHILDREN_NUM),
	}
}
