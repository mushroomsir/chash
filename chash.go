package chash

import (
	"hash/fnv"
	"sort"
	"strconv"
	"sync"
)

// New ...
func New(virtual int) *Nodes {
	node := &Nodes{
		nodes:         make(map[uint64]string),
		orderedCircle: make([]uint64, 0),
		virtualNode:   virtual,
	}
	if node.virtualNode < 1 {
		node.virtualNode = 1
	}
	return node
}

// Nodes ...
type Nodes struct {
	nodes         map[uint64]string
	orderedCircle []uint64
	virtualNode   int
	lock          sync.RWMutex
}

// AddNode ...
func (n *Nodes) AddNode(nodes ...string) {
	if len(nodes) == 0 {
		return
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	for _, realNode := range nodes {
		for i := 0; i < n.virtualNode; i++ {
			virtualNode := realNode + strconv.Itoa(i)
			hashcode := hash(virtualNode)
			n.nodes[hashcode] = realNode
			n.orderedCircle = append(n.orderedCircle, hashcode)
		}
	}
	sort.Slice(n.orderedCircle, func(i, j int) bool { return n.orderedCircle[i] < n.orderedCircle[j] })
}

// TragetNode ...
func (n *Nodes) TragetNode(key string) string {
	n.lock.RLock()
	defer n.lock.RUnlock()
	node := n.nodes[isRange(hash(key), n.orderedCircle)]
	return node
}
func isRange(hashcode uint64, circle []uint64) uint64 {
	if hashcode > circle[len(circle)-1] && circle[0] >= hashcode {
		return circle[0]
	}
	for i, val := range circle {
		if val >= hashcode {
			return circle[i]
		}
	}
	return circle[0]
}
func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
