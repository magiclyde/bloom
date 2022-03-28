/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2022/3/28 下午4:00
 * @note: bloom filter in go
 * @refer: https://codeburst.io/lets-implement-a-bloom-filter-in-go-b2da8a4b849f
 */

package bloom

import (
	"github.com/magiclyde/murmur3"
	"hash"
	"hash/fnv"
)

type Interface interface {
	Add(item []byte)       // Adds the item into the Set
	Test(item []byte) bool // Check if items is maybe in the Set
}

// BloomFilter probabilistic data structure definition
type BloomFilter struct {
	bitset    []bool        // The bloom-filter bitset
	k         uint          // Number of hash values
	n         uint          // Number of elements in the filter
	m         uint          // Size of the bloom filter bitset
	hashFuncs []hash.Hash64 // The hash functions
}

// New Returns a new BloomFilter object,
func New(size uint) *BloomFilter {
	return &BloomFilter{
		bitset:    make([]bool, size),
		k:         3, // k = ln(2) * m / n
		m:         size,
		n:         uint(0),
		hashFuncs: []hash.Hash64{murmur3.New64(), fnv.New64(), fnv.New64a()},
	}
}

// Add the item into the bloom filter set by hashing in over the hash functions
func (bf *BloomFilter) Add(item []byte) {
	hashes := bf.hashValues(item)
	i := uint(0)
	for {
		if i >= bf.k {
			break
		}
		position := uint(hashes[i]) % bf.m
		bf.bitset[uint(position)] = true
		i += 1
	}
	bf.n += 1
}

// Calculates all the hash values by applying in the item over the hash functions
func (bf *BloomFilter) hashValues(item []byte) []uint64 {
	var result []uint64
	for _, hashFunc := range bf.hashFuncs {
		hashFunc.Write(item)
		result = append(result, hashFunc.Sum64())
		hashFunc.Reset()
	}
	return result
}

// Test if the item into the bloom filter is set by hashing in over // the hash functions
func (bf *BloomFilter) Test(item []byte) (exists bool) {
	hashes := bf.hashValues(item)
	i := uint(0)
	exists = true

	for {
		if i >= bf.k {
			break
		}

		position := uint(hashes[i]) % bf.m
		if !bf.bitset[uint(position)] {
			exists = false
			break
		}
		i += 1
	}
	return
}
