package prefixtree

type Walker interface {
	Walk(data []byte) (node Node, err bool)
}

type walker struct {
	currentBranch Branch
	currentIndex  int
}

func NewWalker(root Branch) Walker {
	return &walker{
		currentBranch: root,
	}
}

/*
Walk walks lazily (if no error was returned before, you can continue
matching a string by calling this function one more time with a next
part of data)

Returned node is the last one node. In case of failure, it is unmatched
node (that wanted to be, but dreams aren't always real)
*/
func (w *walker) Walk(data []byte) (node Node, err bool) {
	for _, char := range data {
		node = w.currentBranch[w.currentIndex]

		if node.char != char {
			newBranch := getBranch(node.variants, char)

			if newBranch == nil {
				return node, true
			}

			w.currentBranch = newBranch
			w.currentIndex = 1
		} else {
			w.currentIndex++
		}
	}

	return w.currentBranch[w.currentIndex-1], false
}

func IsTail(node Node) bool {
	for _, branch := range node.variants {
		if branch == nil {
			return true
		}
	}

	return false
}

/*
getBranch returns a branch a first node of which one matches
a char we need. If nil is returned, this means that we could
not find a corresponding branch (tailing branch cannot be
returned)
*/
func getBranch(branches []Branch, char byte) Branch {
	for _, branch := range branches {
		if branch[0].char == char {
			return branch
		}
	}

	return nil
}
