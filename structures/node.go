package structures

type Node struct {
	Child       map[string]*Node
	Frequency   int32
	IsEndOfWord bool
}

func NewNode() *Node {
	node := &Node{
		Child:       nil,
		Frequency:   0,
		IsEndOfWord: false,
	}

	return node
}
