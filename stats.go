// Package stats implements simple statistical functions.
package stats

import (
	"sort"
)

// Quartile ...
func (q Q) Quartile(slice []float64) float64 {

	if q >= 4 || q < 1 {
		panic("q not acceptable")
	}

	var s sort.Float64Slice

	s = slice
	s.Sort()
	l := s.Len()

	r := (l + 1) * int(q) / 4

	return s[r]

}

// InterQuartileRange ...
func InterQuartileRange(slice []float64) float64 {
	var s sort.Float64Slice

	s = slice
	s.Sort()
	l := s.Len()

	q1 := (l + 1) / 4
	q3 := (l + 1) * 3 / 4

	return s[q3] - s[q1]
}

// Percentile ...
func Percentile(slice []float64, p float64) float64 {
	if p > 1 || p < 0 {
		panic("value for p not allowed")
	}
	var s sort.Float64Slice

	s = slice
	s.Sort()
	l := float64(s.Len()) * p

	index := int(l)

	return s[index]
}

// Percentile Method.
func (p P) Percentile(slice []float64) float64 {

	var s sort.Float64Slice

	s = slice
	s.Sort()
	l := float64(s.Len()) * float64(p/100)

	index := int(l)

	return s[index]
}
