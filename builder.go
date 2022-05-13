package prefixtree

const (
	head = byte(0)
	tail = byte(1)
)

type Leaf struct {
	char   byte
	leaves []Leaf
}

func BuildTree(branches [][]byte) Leaf {
	longestBranchIndex := longestArray(branches)
	masterBranch := buildMasterBranch(branches[longestBranchIndex])
	otherBranches := append(branches[:longestBranchIndex], branches[longestBranchIndex+1:]...)

	for _, branch := range otherBranches {
		masterBranch = addBranch(masterBranch, branch)
	}

	return masterBranch
}

func InsertOne(base Leaf, branch []byte) Leaf {
	return addBranch(base, branch)
}

func InsertMany(base Leaf, branches [][]byte) Leaf {
	for _, branch := range branches {
		base = addBranch(base, branch)
	}

	return base
}

func newLeaf(char byte) *Leaf {
	return &Leaf{
		char:   char,
		leaves: make([]Leaf, 0, 1),
	}
}

func newTailLeaf() *Leaf {
	return &Leaf{
		char:   tail,
		leaves: nil,
	}
}

func (l Leaf) IsTail() bool {
	return getLeaf(l.leaves, tail) != nil
}

/*
buildMasterBranch builds a root leaf from zero
*/
func buildMasterBranch(branch []byte) (leaf Leaf) {
	if len(branch) == 0 {
		return leaf
	}

	rootLeaf := &Leaf{
		char:   head,
		leaves: make([]Leaf, 0, 1),
	}

	currentLeaf := rootLeaf

	for _, char := range branch {
		leaf = *newLeaf(char)
		currentLeaf.leaves = append(currentLeaf.leaves, leaf)
		currentLeaf = &currentLeaf.leaves[len(currentLeaf.leaves)-1]
	}

	currentLeaf.leaves = append(currentLeaf.leaves, *newTailLeaf())

	return *rootLeaf
}

func addBranch(base Leaf, branch []byte) Leaf {
	commonPrefixOffset := -1
	currentLeaf := &base

	for i, char := range branch {
		leaf := getLeaf(currentLeaf.leaves, char)

		if leaf == nil {
			commonPrefixOffset = i
			break
		}

		currentLeaf = leaf
	}

	if commonPrefixOffset == -1 {
		if getLeaf(currentLeaf.leaves, tail) == nil {
			currentLeaf.leaves = append(currentLeaf.leaves, *newTailLeaf())
		}

		return base
	}

	for _, char := range branch[commonPrefixOffset:] {
		leaf := *newLeaf(char)
		currentLeaf.leaves = append(currentLeaf.leaves, leaf)
		currentLeaf = &currentLeaf.leaves[len(currentLeaf.leaves)-1]
	}

	currentLeaf.leaves = append(currentLeaf.leaves, *newTailLeaf())

	return base
}

func getLeaf(leaves []Leaf, char byte) *Leaf {
	for i := 0; i < len(leaves); i++ {
		if leaves[i].char != char {
			continue
		}

		return &leaves[i]
	}

	return nil
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
