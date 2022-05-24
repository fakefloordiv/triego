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

func testTree(t *testing.T, rootNode Node, inputBranch []byte) {
	t.Run("Part of correct input data", func(t *testing.T) {
		walker := NewWalker(rootNode)
		firstHalf := inputBranch[:len(inputBranch)/2]
		Node, err := walker.Walk(firstHalf)

		if err {
			t.Error("test failed: no match found")
		} else if Node.IsTail() {
			t.Error("returned Node must not be tailing")
		}
	})

	t.Run("Two parts of correct data", func(t *testing.T) {
		walker := NewWalker(rootNode)
		firstHalf, secondHalf := inputBranch[:len(inputBranch)/2], inputBranch[len(inputBranch)/2:]
		Node, err := walker.Walk(firstHalf)

		if err {
			t.Error("test failed: no match found for first half")
		} else if Node.IsTail() {
			t.Error("returned Node must not be tailing")
		}

		Node, err = walker.Walk(secondHalf)

		if err {
			t.Error("test failed: no match found")
		} else if !Node.IsTail() {
			t.Error("returned Node must be tailing")
		}
	})

	t.Run("Full correct input data", func(t *testing.T) {
		walker := NewWalker(rootNode)
		Node, err := walker.Walk(inputBranch)

		if err {
			t.Error("test failed: no match found")
		} else if !Node.IsTail() {
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
	rootNode := BuildTree(multipleBranchesShort)
	rootNode = InsertMany(rootNode, insertMany)
	testTree(t, rootNode, tryToMatch)
}
