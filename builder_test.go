package prefixtree

import (
	"strings"
	"testing"
)

var (
	singleBranch = [][]byte{
		[]byte("Hello"),
	}

	multipleBranchesShort = [][]byte{
		[]byte(strings.Repeat("abc", 6)),
		[]byte(strings.Repeat("def", 6)),
		[]byte(strings.Repeat("jkl", 6)),
	}

	multipleBranchesLong = [][]byte{
		[]byte(strings.Repeat("abc", 100)),
		[]byte(strings.Repeat("def", 100)),
		[]byte(strings.Repeat("jkl", 100)),
	}
)

func testTree(t *testing.T, tree Branch, inputBranch []byte) {
	t.Run("Part of correct input data", func(t *testing.T) {
		myWalker := NewWalker(tree)
		firstHalf := inputBranch[:len(inputBranch)/2]
		node, err := myWalker.Walk(firstHalf)

		if err {
			t.Error("test failed: no match found")
		} else if IsTail(node) {
			t.Error("returned Node must not be tailing")
		}
	})

	t.Run("Two parts of correct data", func(t *testing.T) {
		myWalker := NewWalker(tree)
		firstHalf, secondHalf := inputBranch[:len(inputBranch)/2], inputBranch[len(inputBranch)/2:]
		node, err := myWalker.Walk(firstHalf)

		if err {
			t.Error("test failed: no match found for first half")
		} else if IsTail(node) {
			t.Error("returned Node must not be tailing")
		}

		node, err = myWalker.Walk(secondHalf)

		if err {
			t.Error("test failed: no match found")
		} else if !IsTail(node) {
			t.Error("returned Node must be tailing", node)
		}
	})

	t.Run("Full correct input data", func(t *testing.T) {
		myWalker := NewWalker(tree)
		node, err := myWalker.Walk(inputBranch)

		if err {
			t.Error("test failed: no match found")
		} else if !IsTail(node) {
			t.Error("returned Node must be tailing")
		}
	})
}

func TestSingleBranchBuild(t *testing.T) {
	tryToMatch := []byte("Hello")
	rootNode := BuildTree(singleBranch)
	testTree(t, rootNode, tryToMatch)
}

func TestMultipleBranchesBuild(t *testing.T) {
	tryToMatch := []byte(strings.Repeat("def", 6))
	rootNode := BuildTree(multipleBranchesShort)
	testTree(t, rootNode, tryToMatch)
}

func TestSingleBranchInsertOne(t *testing.T) {
	insertOne := []byte("world")
	rootNode := BuildTree(singleBranch)
	rootNode = InsertOne(rootNode, insertOne)
	testTree(t, rootNode, insertOne)
}

func TestMultipleBranchesInsertOne(t *testing.T) {
	insertOne := []byte(strings.Repeat("kol", 6))
	rootNode := BuildTree(multipleBranchesShort)
	rootNode = InsertOne(rootNode, insertOne)
	testTree(t, rootNode, insertOne)
}

func TestSingleBranchInsertMany(t *testing.T) {
	tryToMatch := []byte("lorem ipsum")
	insertMany := [][]byte{
		[]byte("world"),
		tryToMatch,
	}
	rootNode := BuildTree(singleBranch)
	rootNode = InsertMany(rootNode, insertMany)
	testTree(t, rootNode, tryToMatch)
}

func TestMultipleBranchesInsertMany(t *testing.T) {
	tryToMatch := []byte(strings.Repeat("nmb", 6))
	insertMany := [][]byte{
		[]byte(strings.Repeat("kol", 6)),
		tryToMatch,
	}
	rootBranch := BuildTree(multipleBranchesShort)
	rootBranch = InsertMany(rootBranch, insertMany)

	testTree(t, rootBranch, tryToMatch)
}
