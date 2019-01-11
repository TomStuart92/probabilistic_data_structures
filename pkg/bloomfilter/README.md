# Bloom Filters

A Bloom filter is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set". (https://en.wikipedia.org/wiki/Bloom_filter)

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