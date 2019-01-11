# Count Min Sketch

In computing, the count–min sketch (CM sketch) is a probabilistic data structure that serves as a frequency table of events in a stream of data. It uses hash functions to map events to frequencies, but unlike a hash table uses only sub-linear space, at the expense of overcounting some events due to collisions. (https://en.wikipedia.org/wiki/Count%E2%80%93min_sketch)

## API

```go
capacity := 1000  // total data set size
errRate := 0.001  // acceptable error rate i.e. we should have no more than 0.1% false positive matches

bloomFilter := NewBloomFilter(capacity, errRate)
bloomFilter.add("test");
bloomFilter.mayContain("test")  // true (hopefully)
```

## Dependencies

We use the murmur3 hashing algorithm (https://github.com/spaolacci/murmur3) and the bitset library (https://github.com/willf/bitset) to track flipped bits.