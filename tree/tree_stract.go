package tree

// Node 树中节点的结构（The structure of the nodes in the tree）
type Node struct {
	// Data 数据信息（Data）
	Data interface{}

	// LeftChildPtr 指向最左边子节点的指针（Pointer to the leftmost child node）
	LeftChildPtr *Node

	// RightChildPtr 指向最右边子节点的指针（Pointer to the rightmost child node）
	RightChildPtr *Node

	// ChildNodePtrs 指向所有子节点的指针数组，index=0为最左子节点，index=n为最右子节点（Array of pointers to all child nodes, index=0 is the leftmost child node, index=n is the rightmost child node）
	ChildNodePtrs []*Node

	// ParentNodePtr 指向当前节点父节点的指针（Pointer to the parent node of the current node）
	ParentNodePtr *Node

	// Dynasty 当前节点代，Root节点为0（Current node generation, Root node is 0）
	Dynasty int
}

// Tree 树的结构（The structure of the Tree）
type Tree struct {
	// Root 树的根节点（Root node of tree）
	Root *Node

	// NodesPtrs 树的所有节点指针集合，顺序按插入顺序（The set of all node pointers of the tree, in the order of insertion）
	NodesPtrs []*Node

	// Degree 树的度（Degree for tree）
	Degree int

	// insCur 插入游标，指向最后一个被插入的节点的指针（Insert cursor, pointer to the last inserted node）
	insCur *Node
}
