package prefixtree

const (
	head = byte(0)
	tail = byte(1)
)

type Node struct {
	char   byte
	leaves []Node
}

func BuildTree(branches [][]byte) Node {
	longestBranchIndex := longestArray(branches)
	masterBranch := buildMasterBranch(branches[longestBranchIndex])
	otherBranches := append(branches[:longestBranchIndex], branches[longestBranchIndex+1:]...)

	for _, branch := range otherBranches {
		masterBranch = addBranch(masterBranch, branch)
	}

	return masterBranch
}

func InsertOne(base Node, branch []byte) Node {
	return addBranch(base, branch)
}

func InsertMany(base Node, branches [][]byte) Node {
	for _, branch := range branches {
		base = addBranch(base, branch)
	}

	return base
}

func newNode(char byte) *Node {
	return &Node{
		char:   char,
		leaves: make([]Node, 0, 1),
	}
}

func newTailNode() *Node {
	return &Node{
		char:   tail,
		leaves: nil,
	}
}

func (l Node) IsTail() bool {
	return getNode(l.leaves, tail) != nil
}

/*
buildMasterBranch builds a root Node from zero
*/
func buildMasterBranch(branch []byte) (node Node) {
	if len(branch) == 0 {
		return node
	}

	rootNode := &Node{
		char:   head,
		leaves: make([]Node, 0, 1),
	}

	currentNode := rootNode

	for _, char := range branch {
		node = *newNode(char)
		currentNode.leaves = append(currentNode.leaves, node)
		currentNode = &currentNode.leaves[len(currentNode.leaves)-1]
	}

	currentNode.leaves = append(currentNode.leaves, *newTailNode())

	return *rootNode
}

func addBranch(base Node, branch []byte) Node {
	commonPrefixOffset := -1
	currentNode := &base

	for i, char := range branch {
		Node := getNode(currentNode.leaves, char)

		if Node == nil {
			commonPrefixOffset = i
			break
		}

		currentNode = Node
	}

	if commonPrefixOffset == -1 {
		if getNode(currentNode.leaves, tail) == nil {
			currentNode.leaves = append(currentNode.leaves, *newTailNode())
		}

		return base
	}

	for _, char := range branch[commonPrefixOffset:] {
		Node := *newNode(char)
		currentNode.leaves = append(currentNode.leaves, Node)
		currentNode = &currentNode.leaves[len(currentNode.leaves)-1]
	}

	currentNode.leaves = append(currentNode.leaves, *newTailNode())

	return base
}

func getNode(leaves []Node, char byte) *Node {
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
