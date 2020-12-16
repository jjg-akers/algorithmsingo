package datastructures

type Tree struct {
	root *node
}

func (t *Tree) Add(val int) *Tree {
	if t.root == nil {
		t.root = &node{
			val:   val,
			left:  nil,
			right: nil,
		}
	} else {
		t.root.add(val)
	}

	return t
}

type node struct {
	val   int
	left  *node
	right *node
}

func (n *node) add(val int) {
	if n == nil {
		return
	} else if val <= n.val {
		if n.left == nil {
			n.left = &node{
				val:   val,
				left:  nil,
				right: nil,
			}
		} else {
			n.left.add(val)
		}
	} else {
		if n.right == nil {
			n.right = &node{
				val:   val,
				left:  nil,
				right: nil,
			}
		} else {
			n.right.add(val)
		}
	}
}

func (t *Tree) DFS(val int) *node {
	if t == nil {
		return nil
	}

	if t.root.val == val {
		return t.root
	}

	return t.root.dfs(val)
}

func (n *node) dfs(val int) *node {
	if n == nil {
		return nil
	}

	if n.val == val {
		return n
	}

	return n.dfs(val)
}
