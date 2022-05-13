package prefixtree

type Walker interface {
	Walk(data []byte) (leaf Leaf, err bool)
}

type walker struct {
	currentLeaf Leaf
}

func NewWalker(root Leaf) Walker {
	return &walker{
		currentLeaf: root,
	}
}

func (w *walker) Walk(data []byte) (leaf Leaf, err bool) {
	for _, char := range data {
		childLeaf := getLeaf(w.currentLeaf.leaves, char)

		if childLeaf == nil {
			return w.currentLeaf, true
		}

		w.currentLeaf = *childLeaf
	}

	return w.currentLeaf, false
}
