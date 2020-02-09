package tree

// Tree holds elements of the red-black tree
type Tree struct {
	Root       *Node
	size       int
}

// Node is a single element within the tree
type Node struct {
	Key       interface{}
	Value     interface{}
	nodeIndex int
	Left      *Node
	Right     *Node
	Parent    *Node
}