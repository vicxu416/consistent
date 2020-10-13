package consistent

import (
	"fmt"

	"github.com/vicxu416/consistent.git/hashing"
)

func DefaultConfig() *Config {
	return &Config{
		numberOfReplicas: 3,
		hashFunc:         hashing.CRC32Hashing,
		nodeIDFunc:       defaultNodeIDGen,
	}
}

func defaultNodeIDGen(nodeID string, replicNo int) string {
	return fmt.Sprintf("%s-%d", nodeID, replicNo)
}

type Option func(config *Config)

func SetHashing(hashTye hashing.HashType) Option {
	return func(config *Config) {
		switch hashTye {
		case hashing.CRC32:
			config.hashFunc = hashing.CRC32Hashing
		case hashing.FNV:
			config.hashFunc = hashing.FNVHashing
		case hashing.MURMUR:
			config.hashFunc = hashing.MurMurHashing
		default:
			config.hashFunc = hashing.CRC32Hashing
		}
	}
}

func SetReplicas(num int64) Option {
	return func(config *Config) {
		config.numberOfReplicas = num
	}
}

func SetNodeIDGen(fn NodeIDGenerator) Option {
	return func(config *Config) {
		config.nodeIDFunc = fn
	}
}
