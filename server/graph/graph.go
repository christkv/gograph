package graph

import "sync/atomic"

type Graph struct {
	id        uint64	// Current graph unique id
	nodeId    uint64	// Next node id
	name      string	// Name of graph
	entryNode *Node		// Initial node entry
}

func (p *Graph) init() {
	nodeId := atomic.AddUint64(&p.nodeId, uint64(1))
	p.entryNode = CreateNode(nodeId)
}
