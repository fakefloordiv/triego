package prefixtree

type Walker interface {
	Walk(data []byte) (leaf Node, err bool)
}

type walker struct {
	currentNode Node
}

func NewWalker(root Node) Walker {
	return &walker{
		currentNode: root,
	}
}

func (w *walker) Walk(data []byte) (node Node, err bool) {
	for _, char := range data {
		childNode := getNode(w.currentNode.leaves, char)

		if childNode == nil {
			return w.currentNode, true
		}

		w.currentNode = *childNode
	}

	return w.currentNode, false
}
