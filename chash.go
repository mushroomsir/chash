package chash

import (
	"hash/fnv"
	"sort"
)

// Nodes ...
type Nodes struct {
	nodes         map[uint64]string
	orderedCircle []uint64
}

// AddNode ...
func (n *Nodes) AddNode(nodes ...string) {
	if len(nodes) == 0 {
		return
	}
	for _, v := range nodes {
		hashcode := hash(v)
		n.nodes[hash(v)] = v
		n.orderedCircle = append(n.orderedCircle, hashcode)
	}
	sort.Slice(n.orderedCircle, func(i, j int) bool { return n.orderedCircle[i] < n.orderedCircle[j] })
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
