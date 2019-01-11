package countminsketch

import (
	"log"
	"os"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := New(7, 2000)
	if err != nil {
		t.Error(err)
	}
	c.UpdateString("test", 1)
	if c.EstimateString("test") != 1 {
		t.Error("Inaccurate Estimate")
	}
}

func TestAccuracy(t *testing.T) {
	log.SetOutput(os.Stdout)
	s, err := NewWithEstimates(0.0001, 0.9999)
	if err != nil {
		t.Error(err)
	}

	iterations := 5500
	var diverged int
	for i := 1; i < iterations; i++ {
		v := uint64(i % 50)

		s.UpdateString(strconv.Itoa(i), v)
		vv := s.EstimateString(strconv.Itoa(i))
		if vv > v {
			diverged++
		}
	}

	var miss int
	for i := 1; i < iterations; i++ {
		vv := uint64(i % 50)

		v := s.EstimateString(strconv.Itoa(i))
		if v != vv {
			log.Printf("real: %d, estimate: %d\n", vv, v)
			miss++
		}
	}
	log.Printf("missed %d of %d (%d diverged during adds)", miss, iterations, diverged)
}
