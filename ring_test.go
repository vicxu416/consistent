package consistent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nodes = []string{"192.168.0.1", "192.168.0.2", "192.168.0.3", "192.168.0.4", "192.168.0.5"}

func TestAddNode(t *testing.T) {
	ring := NewDefault()

	err := ring.AddNodes(nodes...)
	assert.NoError(t, err)
	assert.Len(t, ring.Nodes(), len(nodes))
}

func TestGetNode(t *testing.T) {
	ring := NewDefault()

	err := ring.AddNodes(nodes...)
	assert.NoError(t, err)
	key := "test123"
	node := ring.GetNode(key)
	exist := false
	for _, n := range nodes {
		if n == node {
			exist = true
			break
		}
	}
	assert.True(t, exist)
}

func TestRemoveNode(t *testing.T) {
	ring := NewDefault()

	err := ring.AddNodes(nodes...)
	assert.NoError(t, err)

	ring.RemoveNodes(nodes[0])
	assert.Len(t, ring.Nodes(), len(nodes)-1)
}
