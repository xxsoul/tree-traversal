package tree

// AddNode 向树中添加节点，如果树中当前节点数为0，则添加到根节点上；否则按着广度优先的方式挂到下面节点中。
func (tree *Tree) AddNode(data interface{}) {
	insCur := tree.insCur
	insCurParent := tree.selectParentToAdd()
	if insCurParent == nil {
		insCurParent = insCur
	}

	node := new(Node)
	tree.insCur = node

	tree.insCur.Data = data
	tree.insCur.ParentNodePtr = insCurParent
	tree.insCur.ChildNodePtrs = make([]*Node, 0)
	if insCurParent != nil {
		tree.insCur.Dynasty = insCurParent.Dynasty + 1
	}

	if tree.Root == nil {
		// 初始化的情况，root为nil（In the case of initialization, root is nil）
		tree.Root = node
		tree.insCur = node
		tree.NodesPtrs = make([]*Node, 0)
	} else {
		// 非初始化的情况，先插入左子节点，然后右子节点
		insCurParent.ChildNodePtrs = append(insCurParent.ChildNodePtrs, node)
		if len(insCurParent.ChildNodePtrs) == 1 {
			insCurParent.LeftChildPtr = node
			insCurParent.RightChildPtr = node
		} else {
			insCurParent.RightChildPtr = insCurParent.ChildNodePtrs[len(insCurParent.ChildNodePtrs)-1]
		}
	}

	tree.NodesPtrs = append(tree.NodesPtrs, tree.insCur)
}

// selectParentToAdd 为节点选择一个父节点
func (tree *Tree) selectParentToAdd() *Node {
	insCur := tree.insCur
	if insCur == nil {
		return insCur
	}
	parent := insCur.ParentNodePtr

	if parent == nil {
		return insCur
	}
	if len(parent.ChildNodePtrs) < tree.Degree {
		return parent
	}

	// 由于插入时广度优先的，所以只要顺序遍历指针数组，得到的就是从Root广度优先到的扫描结果
	for _, n := range tree.NodesPtrs {
		if len(n.ChildNodePtrs) < tree.Degree {
			return n
		}
	}
	return nil
}
