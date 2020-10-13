package consistent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	testcase := []uint32{1, 2, 3, 5, 6, 7, 10, 11}

	keys := SortedKeys([]HashKey{})
	keys.Insert(hashKey(7))
	keys.Insert(hashKey(5))
	keys.Insert(hashKey(3))
	keys.Insert(hashKey(1), hashKey(2))
	keys.Insert(hashKey(6))
	keys.Insert(hashKey(11), hashKey(10))

	assert.Len(t, keys, len(testcase))
	for i := range testcase {
		assert.Equal(t, keys[i].Val(), testcase[i])
	}
}

func TestDel(t *testing.T) {
	testcase := []uint32{1, 2, 3, 5, 7}
	keys := SortedKeys([]HashKey{})
	keys.Insert(hashKey(7), hashKey(1), hashKey(2), hashKey(3), hashKey(5))
	keys.Insert(hashKey(12), hashKey(10), hashKey(11))

	keys.Del(hashKey(10), hashKey(12), hashKey(11), hashKey(13))

	for i := range keys {
		assert.Equal(t, keys[i].Val(), testcase[i])
	}
}

func TestFind(t *testing.T) {
	keys := SortedKeys([]HashKey{})
	keys.Insert(hashKey(7), hashKey(1), hashKey(5), hashKey(2))

	assert.Equal(t, keys.Find(hashKey(4)).Val(), uint32(5))
	assert.Equal(t, keys.Find(hashKey(8)).Val(), uint32(1))
	assert.Equal(t, keys.Find(hashKey(2)).Val(), uint32(2))
	assert.Equal(t, keys.Find(hashKey(3)).Val(), uint32(5))
	assert.Equal(t, keys.Find(hashKey(6)).Val(), uint32(7))
}
