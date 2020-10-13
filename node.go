package consistent

func newNodes() *Nodes {
	return &Nodes{
		sortedKeys: SortedKeys(make([]HashKey, 0, 1)),
		nodeKeyMap: make(map[string][]HashKey),
	}
}

type Nodes struct {
	sortedKeys SortedKeys
	nodeKeyMap map[string][]HashKey
}

func (nodes *Nodes) Exist(nodeID string) bool {
	keys, exist := nodes.nodeKeyMap[nodeID]
	return exist && len(keys) > 0
}

func (nodes *Nodes) Find(key HashKey) HashKey {
	return nodes.sortedKeys.Find(key)
}

func (nodes *Nodes) AllNodes() []string {
	nodeIDs := make([]string, 0, len(nodes.nodeKeyMap))

	for nodeID := range nodes.nodeKeyMap {
		nodeIDs = append(nodeIDs, nodeID)
	}
	return nodeIDs
}

func (nodes *Nodes) Add(nodeID string, keys ...HashKey) error {
	nodes.sortedKeys.Insert(keys...)
	nodes.nodeKeyMap[nodeID] = keys
	return nil
}

func (nodes *Nodes) Del(nodeID string) ([]HashKey, error) {
	keys, ok := nodes.nodeKeyMap[nodeID]
	if !ok {
		return nil, nil
	}
	delete(nodes.nodeKeyMap, nodeID)
	nodes.sortedKeys.Del(keys...)
	return keys, nil
}
