package consistent

import (
	"sync"
)

type NodeIDGenerator func(nodeID string, replicasNo int) string

type Config struct {
	numberOfReplicas int64
	hashFunc         HashFunc
	nodeIDFunc       NodeIDGenerator
}

func New(options ...Option) *HashRing {
	config := DefaultConfig()

	for _, opt := range options {
		opt(config)
	}

	return &HashRing{
		Config:      config,
		nodeHashMap: make(map[uint32]string),
		nodes:       newNodes(),
	}

}

func NewDefault() *HashRing {
	config := DefaultConfig()

	return &HashRing{
		Config:      config,
		nodeHashMap: make(map[uint32]string),
		nodes:       newNodes(),
	}
}

type HashRing struct {
	sync.RWMutex

	nodes *Nodes
	*Config
	nodeHashMap map[uint32]string
}

func (ring *HashRing) GetNode(key string) string {
	ring.RLock()
	defer ring.RUnlock()
	hashKey := ring.genKey(key)
	gotKey := ring.nodes.Find(hashKey)
	if gotKey == nil {
		return ""
	}
	nodeID, exist := ring.nodeHashMap[gotKey.Val()]
	if !exist {
		return ""
	}

	return nodeID
}

func (ring *HashRing) ResetNodes(nodeIDs ...string) error {
	ring.Lock()
	defer ring.Unlock()

	currNodeIDs := ring.Nodes()
	expectNodes := make(map[string]bool)
	for i := range nodeIDs {
		expectNodes[nodeIDs[i]] = true
	}

	for _, nodeID := range currNodeIDs {
		if !expectNodes[nodeID] {
			ring.remove(nodeID)
		}
	}

	for _, nodeID := range nodeIDs {
		if !ring.nodes.Exist(nodeID) {
			ring.addNode(nodeID)
		}
	}

	return nil
}

func (ring *HashRing) Nodes() []string {
	ring.RLock()
	defer ring.RUnlock()
	return ring.nodes.AllNodes()
}

func (ring *HashRing) AddNodes(nodeIDs ...string) error {
	ring.Lock()
	defer ring.Unlock()

	for _, nodeID := range nodeIDs {
		ring.addNode(nodeID)
	}
	return nil
}

func (ring *HashRing) RemoveNodes(nodeIDs ...string) error {
	ring.Lock()
	defer ring.Unlock()

	for _, nodeID := range nodeIDs {
		ring.remove(nodeID)
	}
	return nil
}

func (ring *HashRing) remove(nodeID string) error {
	keys, err := ring.nodes.Del(nodeID)
	if err != nil {
		return err
	}

	for i := range keys {
		delete(ring.nodeHashMap, keys[i].Val())
	}
	return nil
}

func (ring *HashRing) addNode(nodeID string) error {
	hashKeys := make([]HashKey, ring.numberOfReplicas)
	for i := range hashKeys {
		nodeIndexID := ring.nodeIDFunc(nodeID, i)
		hashKeys[i] = ring.genKey(nodeIndexID)
		ring.nodeHashMap[hashKeys[i].Val()] = nodeID
	}
	ring.nodes.Add(nodeID, hashKeys...)
	return nil
}

func (ring *HashRing) genKey(nodeID string) HashKey {
	return hashKey(ring.hashFunc(nodeID))
}
