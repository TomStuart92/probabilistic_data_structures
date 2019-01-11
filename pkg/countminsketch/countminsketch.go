package countminsketch

import (
	"errors"
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

type CountMinSketch struct {
	d      uint
	w      uint
	count  [][]uint64
	hasher murmur3.Hash128
}

func New(d uint, w uint) (s *CountMinSketch, err error) {
	if d <= 0 || w <= 0 {
		return nil, errors.New("countminsketch: values of d and w should both be greater than 0")
	}
	s = &CountMinSketch{
		d:      d,
		w:      w,
		count:  make([][]uint64, d),
		hasher: murmur3.New128(),
	}

	for r := uint(0); r < d; r++ {
		s.count[r] = make([]uint64, w)
	}
	return s, nil
}

func NewWithEstimates(epsilon, delta float64) (*CountMinSketch, error) {
	if epsilon <= 0 || epsilon >= 1 {
		return nil, errors.New("countminsketch: value of epsilon should be in range of (0, 1)")
	}
	if delta <= 0 || delta >= 1 {
		return nil, errors.New("countminsketch: value of delta should be in range of (0, 1)")
	}

	w := uint(math.Ceil(2 / epsilon))
	d := uint(math.Ceil(math.Log(1-delta) / math.Log(0.5)))
	// fmt.Printf("ε: %f, δ: %f -> d: %d, w: %d\n", epsilon, delta, d, w)
	return New(d, w)
}

// baseHashes calculates the four base sum hashes based on the origin data
func (s *CountMinSketch) baseHashes(key []byte) (a uint64, b uint64) {
	s.hasher.Reset()
	_, err := s.hasher.Write(key)
	if err != nil {
		fmt.Println("Error Calculating Hash")
	}
	v1, v2 := s.hasher.Sum128()
	return v1, v2
}

func (s *CountMinSketch) locations(key []byte) (locs []uint) {
	locs = make([]uint, s.d)
	a, b := s.baseHashes(key)
	ua := uint(a)
	ub := uint(b)
	for r := uint(0); r < s.d; r++ {
		locs[r] = (ua + ub*r) % s.w
	}
	return
}
func (s *CountMinSketch) Update(key []byte, count uint64) {
	for r, c := range s.locations(key) {
		s.count[r][c] += count
	}
}

func (s *CountMinSketch) UpdateString(key string, count uint64) {
	s.Update([]byte(key), count)
}

// Estimate the frequency of a key. It is point query.
func (s *CountMinSketch) Estimate(key []byte) uint64 {
	var min uint64
	for r, c := range s.locations(key) {
		if r == 0 || s.count[r][c] < min {
			min = s.count[r][c]
		}
	}
	return min
}

// EstimateString estimate the frequency of a key of string
func (s *CountMinSketch) EstimateString(key string) uint64 {
	return s.Estimate([]byte(key))
}
