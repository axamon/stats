package stats_test

import (
	"fmt"

	"github.com/axamon/stats"
)

func ExampleQ_Quartile() {

	slice := []float64{12, 2, 3, 10, 5, 7, 8, 9, 9, 9, 3, 1}

	p1 := stats.Q1.Find(slice)
	p2 := stats.Q2.Find(slice)
	p3 := stats.Q3.Find(slice)

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	// Output:
	// 3
	// 8
	// 9
}

func ExampleInterQuartileRange() {
	slice := []float64{1, 2, 3, 3, 5, 7, 8, 9, 9, 9, 10, 12}

	i := stats.InterQuartileRange(slice)

	fmt.Println(i)
	// Output:
	// 6
}

func ExamplePercentile() {
	slice := []float64{1, 2, 3, 3, 5, 7, 8, 9, 9, 9, 10, 12}

	p01 := stats.FindPercentile(slice, 0.01)
	p10 := stats.FindPercentile(slice, 0.1)
	p50 := stats.FindPercentile(slice, 0.5)
	p90 := stats.FindPercentile(slice, 0.9)
	p99 := stats.FindPercentile(slice, 0.99)

	fmt.Println(p01)
	fmt.Println(p10)
	fmt.Println(p50)
	fmt.Println(p90)
	fmt.Println(p99)
	// Output:
	// 1
	// 2
	// 8
	// 10
	// 12
}

func ExamplePercentile_Find() {
	slice := []float64{12, 2, 8, 3, 9, 7, 3, 9, 5, 9, 10, 1}

	r50 := stats.P50.Find(slice)
	r99 := stats.P99.Find(slice)

	fmt.Println(r50)
	fmt.Println(r99)
	// Output:
	// 8
	// 12
}
