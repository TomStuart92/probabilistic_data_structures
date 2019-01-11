package bloomfilter

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)

/*
A BloomFilter is a probabilistic data structure for tracking members of a set.
It may return false positives i.e. returns true when a member is not a member of the set,
but will never return false negatives.
*/

// BloomFilter main struct
type BloomFilter struct {
	capacity         uint64
	errRate          float64
	numBits          uint64
	numHashFunctions uint64
	byteArray        *bitset.BitSet
}

// calculateNumberOfBits calculates the optimal number of bits to use to achieve the desired errRate false positives
// see derivation at https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
func calculateNumberOfBits(capacity uint64, errRate float64) uint64 {
	return uint64(-float64(capacity) * math.Log(errRate) / (math.Pow(math.Log(2.0), 2.0)))
}

// calculateNumberOfHashFunctions calculates the optimal number of hash functions to use to reduce false positives
// see derivation at https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
func calculateNumberOfHashFunctions(numBits uint64, capacity uint64) uint64 {
	return uint64(float64(numBits/capacity) * math.Log(2.0))
}

// NewBloomFilter creates a new optimised BloomFilter for a given number of items and error rate
func New(capacity uint64, errRate float64) *BloomFilter {
	numBits := calculateNumberOfBits(capacity, errRate)
	numHashFunctions := calculateNumberOfHashFunctions(numBits, capacity)
	return &BloomFilter{
		capacity:         capacity,
		errRate:          errRate,
		numBits:          numBits,
		numHashFunctions: numHashFunctions,
		byteArray:        bitset.New(uint(numBits)),
	}
}

// Add adds a string to the BloomFilter Set by calculating a location, and flipping the bits
func (filter *BloomFilter) Add(s string) *BloomFilter {
	hashes := baseHashes([]byte(s))
	for i := uint64(0); i < filter.numHashFunctions; i++ {
		location := computeKMLocation(hashes, uint64(i))
		byteIndex := uint(location % uint64(filter.numBits)) // take location modulo numBits to ensure we flip bit within an array.
		filter.byteArray.Set(byteIndex)
	}
	return filter
}

/*
Use the Kirsch-Mitzenmacher-Optimization to compute the kth hash.
Using the base hashes h1(x) and h2(x), we can calculate
gi(x) = h1(x)+ i * h2(x), where 0 < i <= k - 1
Allows us to use a set of base hashes to compute all the other needed hashes up to numHashFunctions
See (https://www.eecs.harvard.edu/~michaelm/postscripts/tr-02-05.pdf)
*/
func computeKMLocation(hashes [4]uint64, k uint64) uint64 {
	return hashes[k%2] + k*hashes[2+(((k+(k%2))%4)/2)]
}

/*
MayContain returns true if set may contain element,
but false if it definely does not contain element
*/
func (filter *BloomFilter) MayContain(s string) bool {
	hashes := baseHashes([]byte(s))
	for i := uint64(0); i < filter.numHashFunctions; i++ {
		location := computeKMLocation(hashes, uint64(i))
		byteIndex := uint(location % uint64(filter.numBits))
		if !filter.byteArray.Test(byteIndex) {
			return false
		}
	}
	return true
}

// baseHashes calculates the four base sum hashes based on the origin data
func baseHashes(data []byte) [4]uint64 {
	hasher := murmur3.New128()
	_, err := hasher.Write(data)
	if err != nil {
		fmt.Println("Error Calculating Hash")
	}
	v1, v2 := hasher.Sum128()
	a1 := []byte{1} // spare byte which is appended to get second set of hashes
	_, err = hasher.Write(a1)
	if err != nil {
		fmt.Println("Error Calculating Hash")
	}
	v3, v4 := hasher.Sum128()
	return [4]uint64{
		v1, v2, v3, v4,
	}
}
