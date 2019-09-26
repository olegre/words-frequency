package list

type Trie struct {
	root *Node
	list []*Node
	lnum int
}

func NewTrie(num int) *Trie {
	trie := &Trie{}
	trie.root = NewNode("")
	trie.list = make([]*Node, 0, num)
	trie.lnum = num
	return trie
}

func (trie Trie) GetMostFrequent(i int) []*Node {
	return []*Node{&Node{Letter: "eee", Count: 2}, {Letter: "bbb", Count: 2}}
}

func (trie *Trie) Insert(word string) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		letter := string(word[i])

		var child *Node
		for _, childNode := range node.children {
			if letter == childNode.Letter {
				child = childNode
				node = child
				break
			}
		}
		if child == nil {
			child := NewNode(letter)
			node.children = append(node.children, child)
			node = child
		}

	}
	node.Count++
	node.Word = word

	trie.addToList(node)
}

func (trie *Trie) addToList(node *Node) int {
	index := node.index
	if index < 0 {
		if len(trie.list) < trie.lnum {
			trie.list = append(trie.list, node)
			index = len(trie.list) - 1
		} else if node.Count > trie.list[trie.lnum-1].Count {
			trie.list[trie.lnum-1].index = -1
			trie.list[trie.lnum-1] = node
			index = trie.lnum - 1
		}
	}

	if index > 0 {
		for i := index; i > 0; i-- {
			if trie.list[i].Count > trie.list[i-1].Count {
				trie.list[i-1], trie.list[i] = trie.list[i], trie.list[i-1]
				index = i - 1

			}
		}
	}
	if index < 0 {
		for i := index; i > 0; i-- {
			if trie.list[i].Count > trie.list[i-1].Count {
				trie.list[i-1], trie.list[i] = trie.list[i], trie.list[i-1]
				index = i - 1

			}
		}

	}
	node.index = index
	return index
}

//
//func (Trie) find(letter string, node *Node) (child *Node, err error) {
//	for _, childNode := range node.children {
//		if letter == childNode.Letter {
//			return childNode, nil
//		}
//	}
//	return nil, errors.New("not found")
//}