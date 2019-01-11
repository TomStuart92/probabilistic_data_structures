package main

import "github.com/TomStuart92/probabilistic_data_structures/pkg/bloomfilter"

func main() {
	var capacity uint64 = 1000
	errRate := 0.001
	bloomFilter := bloomfilter.New(capacity, errRate)
	bloomFilter.Add("test")
	bloomFilter.MayContain("test")
}
