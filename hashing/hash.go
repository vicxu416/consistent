package hashing

import (
	"hash/crc32"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type HashType int8

const (
	CRC32 HashType = iota + 1
	FNV
	MURMUR
)

func CRC32Hashing(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func FNVHashing(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func MurMurHashing(key string) uint32 {
	h := murmur3.New32()
	h.Write([]byte(key))
	return h.Sum32()
}
