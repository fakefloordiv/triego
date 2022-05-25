package prefixtree

type Branch []Node

type Node struct {
	char     byte
	variants []Branch
}

func (n Node) GetChar() byte {
	return n.char
}

func (n Node) GetVariants() []Branch {
	return n.variants
}

func BuildTree(branches [][]byte) Branch {
	longestBranchIndex := longestArray(branches)
	masterBranch := buildBranch(branches[longestBranchIndex])
	otherBranches := append(branches[:longestBranchIndex], branches[longestBranchIndex+1:]...)

	for _, branch := range otherBranches {
		masterBranch = addBranch(masterBranch, branch)
	}

	return masterBranch
}

func InsertOne(base Branch, branch []byte) Branch {
	return addBranch(base, branch)
}

func InsertMany(base Branch, branches [][]byte) Branch {
	for _, branch := range branches {
		base = addBranch(base, branch)
	}

	return base
}

/*
buildMasterBranch builds a root Node from zero
*/
func buildBranch(branch []byte) Branch {
	if len(branch) == 0 {
		return Branch{}
	}

	masterBranch := make([]Node, len(branch))

	for index, char := range branch {
		masterBranch[index].char = char
	}

	// long line as hell, but the sense is to add a nil to the
	// last node as nil is an indicator that previous node may
	// be the last one
	masterBranch[len(masterBranch)-1].variants = append(masterBranch[len(masterBranch)-1].variants, nil)

	return masterBranch
}

func addBranch(base Branch, newRawBranch []byte) Branch {
	node, prefix, err := walkByTree(base, newRawBranch)

	if !err {
		// there are 2 cases that may trigger this:
		// 1) new branch is a duplicate and trie already has it
		// 2) new branch is a shorter version of already existing branch
		// So in second case, we need to add a tail marker

		node.variants = append(node.variants, nil)

		return base
	}

	newBranch := buildBranch(newRawBranch[prefix:])
	node.variants = append(node.variants, newBranch)

	return base
}

/*
walkByTree walks by a tree, but used for internals. For example, building
a tree
*/
func walkByTree(root Branch, data []byte) (node *Node, n int, err bool) {
	var (
		index  int
		branch = root
	)

	for _, char := range data {
		node = &branch[index]

		if node.char != char {
			index = 1
			branch = getBranch(node.variants, char)

			if branch == nil {
				return node, n, true
			}
		} else {
			index++
		}

		n++
	}

	return &branch[index-1], n, false
}

func longestArray(arrays [][]byte) (index int) {
	var max int

	for i, arr := range arrays {
		if len(arr) > max {
			max = len(arr)
			index = i
		}
	}

	return index
}
