# TrieGo

### TrieGo is a tiny library for building prefix trees
 Main purpose of the library is to match arrays of bytes, this is why every node keeps its value as a single byte

# How to install?
```bash
$ go get -u github.com/fakefloordiv/triego
```

# How to use?

- `BuildTree(branches [][]byte) Leaf`
  - build a tree with provided branches. Returned `Leaf` struct is a base leaf
- `InsertOne(base Leaf, branch []byte) Leaf`
  - insert a new branch to the tree. Returned `Leaf` is modified base
- `InsertMany(base Leaf, branches [][]byte)`
  - same as just `Insert`, but allows you to insert multiple branches at once
- `NewWalker(root Leaf) Walker` 
  - returns a new walker interface
- `Walker` (interface)
  - `Walk(data []byte) (lastLeaf Leaf, err bool)`
    - walks by a tree, starting with a node last time was the last one. Returns the last one node that was matched and bool that determines whether everything is ok (in case of false, returned leaf is the last one node that was matching an input stream)
- `Leaf` (node struct)
  - `IsTail() bool`
    - determines whether current node may be a finishing node (contains a tail node as a child)
