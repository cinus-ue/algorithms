package binary_tree

type Comparator func(c1 interface{}, c2 interface{}) bool

type BinaryTree struct {
	node       interface{}
	left       *BinaryTree
	right      *BinaryTree
	comparator Comparator
}

func New(com Comparator) *BinaryTree {
	return &BinaryTree{
		node:       nil,
		comparator: com,
	}
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.node == nil {
		return nil
	}
	if tree.node == value {
		return tree
	} else {
		if tree.comparator(value, tree.node) == true {
			t := tree.left.Search(value)
			return t
		} else {
			t := tree.right.Search(value)
			return t
		}
	}

}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.node == nil {
		tree.node = value
		tree.right = New(tree.comparator)
		tree.left = New(tree.comparator)
		return
	} else {
		if tree.comparator(value, tree.node) == true {
			tree.left.Insert(value)
		} else {
			tree.right.Insert(value)
		}
	}

}
