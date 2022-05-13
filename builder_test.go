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

func testTree(t *testing.T, rootLeaf Leaf, inputBranch []byte) {
	t.Run("Part of correct input data", func(t *testing.T) {
		walker := NewWalker(rootLeaf)
		firstHalf := inputBranch[:len(inputBranch)/2]
		leaf, err := walker.Walk(firstHalf)

		if err {
			t.Error("test failed: no match found")
		} else if leaf.IsTail() {
			t.Error("returned leaf must not be tailing")
		}
	})

	t.Run("Two parts of correct data", func(t *testing.T) {
		walker := NewWalker(rootLeaf)
		firstHalf, secondHalf := inputBranch[:len(inputBranch)/2], inputBranch[len(inputBranch)/2:]
		leaf, err := walker.Walk(firstHalf)

		if err {
			t.Error("test failed: no match found for first half")
		} else if leaf.IsTail() {
			t.Error("returned leaf must not be tailing")
		}

		leaf, err = walker.Walk(secondHalf)

		if err {
			t.Error("test failed: no match found")
		} else if !leaf.IsTail() {
			t.Error("returned leaf must be tailing")
		}
	})

	t.Run("Full correct input data", func(t *testing.T) {
		walker := NewWalker(rootLeaf)
		leaf, err := walker.Walk(inputBranch)

		if err {
			t.Error("test failed: no match found")
		} else if !leaf.IsTail() {
			t.Error("returned leaf must be tailing")
		}
	})
}

func TestSingleBranchBuild(t *testing.T) {
	tryToMatch := []byte("Hello")
	rootLeaf := BuildTree(singleBranch)
	testTree(t, rootLeaf, tryToMatch)
}

func TestMultipleBranchesBuild(t *testing.T) {
	tryToMatch := []byte(strings.Repeat("def", 6))
	rootLeaf := BuildTree(multipleBranchesShort)
	testTree(t, rootLeaf, tryToMatch)
}

func TestSingleBranchInsertOne(t *testing.T) {
	insertOne := []byte("world")
	rootLeaf := BuildTree(singleBranch)
	rootLeaf = InsertOne(rootLeaf, insertOne)
	testTree(t, rootLeaf, insertOne)
}

func TestMultipleBranchesInsertOne(t *testing.T) {
	insertOne := []byte(strings.Repeat("kol", 6))
	rootLeaf := BuildTree(multipleBranchesShort)
	rootLeaf = InsertOne(rootLeaf, insertOne)
	testTree(t, rootLeaf, insertOne)
}

func TestSingleBranchInsertMany(t *testing.T) {
	tryToMatch := []byte("lorem ipsum")
	insertMany := [][]byte{
		[]byte("world"),
		tryToMatch,
	}
	rootLeaf := BuildTree(singleBranch)
	rootLeaf = InsertMany(rootLeaf, insertMany)
	testTree(t, rootLeaf, tryToMatch)
}

func TestMultipleBranchesInsertMany(t *testing.T) {
	tryToMatch := []byte(strings.Repeat("nmb", 6))
	insertMany := [][]byte{
		[]byte(strings.Repeat("kol", 6)),
		tryToMatch,
	}
	rootLeaf := BuildTree(multipleBranchesShort)
	rootLeaf = InsertMany(rootLeaf, insertMany)
	testTree(t, rootLeaf, tryToMatch)
}
