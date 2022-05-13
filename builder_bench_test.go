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
			walker := NewWalker(shortTree)
			walker.Walk(shortTreeBranch)
		}
	})

	b.Run("longTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			walker := NewWalker(longTree)
			walker.Walk(longTreeBranch)
		}
	})
}

func BenchmarkManualWalker(b *testing.B) {
	b.Run("shortTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			currentLeaf := shortTree

			for _, char := range shortTreeBranch {
				for _, leaf := range currentLeaf.leaves {
					if leaf.char == char {
						currentLeaf = leaf
						break
					}
				}
			}
		}
	})

	b.Run("longTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			currentLeaf := longTree

			for _, char := range longTreeBranch {
				for _, leaf := range currentLeaf.leaves {
					if leaf.char == char {
						currentLeaf = leaf
						break
					}
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
