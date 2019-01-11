# Count Min Sketch

In computing, the count–min sketch (CM sketch) is a probabilistic data structure that serves as a frequency table of events in a stream of data. It uses hash functions to map events to frequencies, but unlike a hash table uses only sub-linear space, at the expense of overcounting some events due to collisions. (https://en.wikipedia.org/wiki/Count%E2%80%93min_sketch)

## API

```go
c, err := NewWithEstimates(0.0001, 0.9999)
c.UpdateString("test", 1)
c.EstimateString("test")
```

## Dependencies

We use the murmur3 hashing algorithm (https://github.com/spaolacci/murmur3).