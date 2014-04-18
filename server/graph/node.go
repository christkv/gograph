package graph

type Node struct {
	id         uint64
	attributes []*Attribute // Fields values associated with the node
	vertexes   []*Vertex    // List of vertexes for this node
}

func CreateNode(id uint64) *Node {
	return &Node{id: id, attributes: []*Attribute, vertexes: []*Vertex}
}
