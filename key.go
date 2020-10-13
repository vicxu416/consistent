package consistent

import (
	"sort"
)

type SortedKeys []HashKey

func (keys *SortedKeys) Find(key HashKey) HashKey {
	arr := *keys
	if len(arr) == 0 {
		return nil
	}
	i := sort.Search(len(arr), keys.searchFunc(key))
	if i >= keys.Len() {
		return arr[0]
	}
	return arr[i]
}

func (keys *SortedKeys) Del(others ...HashKey) {
	if keys.Len() == 0 {
		return
	}

	for i := range others {
		keys.del(others[i])
	}
}

func (keys *SortedKeys) Insert(others ...HashKey) {
	if cap(*keys) == 0 {
		*keys = make([]HashKey, 0, len(others))
	}

	for i := range others {
		keys.insert(others[i])
	}
}

func (keys SortedKeys) searchFunc(key HashKey) func(i int) bool {
	return func(i int) bool {
		return keys[i].Greater(key) || keys[i].Eq(key)
	}
}

func (keys SortedKeys) Len() int {
	return len(keys)
}

func (keys *SortedKeys) del(key HashKey) SortedKeys {
	arr := *keys
	i := sort.Search(len(arr), keys.searchFunc(key))

	if i >= len(arr) || !arr[i].Eq(key) {
		return *keys
	}
	deletedArr := make([]HashKey, len(arr)-1)
	copy(deletedArr, arr[:i])
	copy(deletedArr[i:], arr[i+1:])
	*keys = deletedArr
	return *keys
}

func (keys *SortedKeys) insert(key HashKey) SortedKeys {
	i := sort.Search(len(*keys), keys.searchFunc(key))
	arr := *keys
	insertedArr := make([]HashKey, len(arr)+1)
	copy(insertedArr, arr[:i])
	copy(insertedArr[i+1:], arr[i:])
	insertedArr[i] = key
	*keys = SortedKeys(insertedArr)
	return insertedArr
}

type HashKey interface {
	Val() uint32
	Less(other HashKey) bool
	Greater(other HashKey) bool
	Eq(other HashKey) bool
}

type hashKey uint32

func (key hashKey) Val() uint32 {
	return uint32(key)
}

func (key hashKey) Less(other HashKey) bool {
	return key.Val() < other.Val()
}

func (key hashKey) Greater(other HashKey) bool {
	return key.Val() > other.Val()
}

func (key hashKey) Eq(other HashKey) bool {
	return key.Val() == other.Val()
}

type HashFunc func(key string) uint32
