package prefixtree

import (
	"strings"
	"testing"
)

var (
	shortTreeBranch = []byte(strings.Repeat("def", 6))
	longTreeBranch  = []byte(strings.Repeat("def", 100))
)

var (
	shortTree = BuildTree(multipleBranchesShort)
	longTree  = BuildTree(multipleBranchesLong)
)

func BenchmarkBuiltInWalker(b *testing.B) {
	b.Run("shortTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			myWalker := NewWalker(shortTree)
			myWalker.Walk(shortTreeBranch)
		}
	})

	b.Run("longTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			myWalker := NewWalker(longTree)
			myWalker.Walk(longTreeBranch)
		}
	})
}

func BenchmarkManualWalker(b *testing.B) {
	b.Run("shortTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var (
				index         int
				currentBranch = shortTree
			)

			for _, char := range shortTreeBranch {
				node := currentBranch[index]

				if node.char != char {
					newBranch := getBranch(node.variants, char)

					currentBranch = newBranch
					index = 1
				} else {
					index++
				}
			}
		}
	})

	b.Run("longTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var (
				index         int
				currentBranch = longTree
			)

			for _, char := range shortTreeBranch {
				node := currentBranch[index]

				if node.char != char {
					newBranch := getBranch(node.variants, char)

					currentBranch = newBranch
					index = 1
				} else {
					index++
				}
			}
		}
	})

	b.Run("hashmap", func(b *testing.B) {
		key := strings.Repeat("def", 100)
		hashmap := map[string][]byte{
			key: []byte("ok"),
		}

		for i := 0; i < b.N; i++ {
			val := hashmap[key]
			noop(val)
		}
	})
}

func noop(_ []byte) {}
