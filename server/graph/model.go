package graph

type Attribute struct {
	Name  string
	Value interface{}
}

type Vertex struct {
	Id         uint64
	Incoming   *Node
	OutGoing   *Node
	Attributes []*Attribute
	Weight     float64
}
