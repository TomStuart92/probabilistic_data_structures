package bloomfilter

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
)

func randStr(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	return str[:len]
}

func TestFalseNegatives(t *testing.T) {
	var capacity uint64 = 1000
	errRate := 0.001
	bloomFilter := NewBloomFilter(capacity, errRate)
	bloomFilter.Add("test")
	if !bloomFilter.MayContain("test") {
		t.Errorf("Expected BloomFilter to not give False Negatives.")
	}
}

func TestFalsePositive(t *testing.T) {
	var capacity uint64 = 1000
	errRate := 0.001
	bloomFilter := NewBloomFilter(capacity, errRate)
	if bloomFilter.MayContain("test") {
		t.Errorf("Expected BloomFilter to not give False Positives.")
	}
}

func TestProbabilitsticFalsePositive(t *testing.T) {
	var capacity uint64 = 1000
	errRate := 0.1
	bloomFilter := NewBloomFilter(capacity, errRate)
	flag := true
	count := 0.0
	for flag {
		count++
		bloomFilter.Add(randStr(10))
		if bloomFilter.MayContain("test") {
			flag = false
		}
		if int(count*errRate) >= 100 {
			t.Errorf("Expected BloomFilter to occasionally give false positives.")
			flag = false
		}
	}

}
