package chash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsRange(t *testing.T) {
	require := require.New(t)
	circle := []uint64{203, 209, 228}

	testcase := []struct {
		val  uint64
		node uint64
	}{
		{val: 192, node: 203},
		{val: 196, node: 203},
		{val: 200, node: 203},
		{val: 204, node: 209},
		{val: 208, node: 209},
		{val: 209, node: 209},
		{val: 212, node: 228},
		{val: 216, node: 228},
		{val: 220, node: 228},
		{val: 228, node: 228},
		{val: 1000, node: 203},
	}
	for _, v := range testcase {
		require.Equal(v.node, isRange(v.val, circle))
	}
}
func TestAddNode(t *testing.T) {
	require := require.New(t)

	nodes := New(100)
	nodes.AddNode("127.0.0.1:1000")
	require.Equal(100, len(nodes.nodes))
	require.Equal(100, len(nodes.orderedCircle))
	nodes.AddNode("127.0.0.1:1001")
	require.Equal(200, len(nodes.nodes))
	require.Equal(200, len(nodes.orderedCircle))

	for i := 0; i < 199; i++ {
		require.True(nodes.orderedCircle[i] < nodes.orderedCircle[i+1])
	}
}
