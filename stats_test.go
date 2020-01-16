package stats_test

import (
	"fmt"

	"github.com/axamon/stats"
)

func ExampleQ_Quartile() {

	slice := []float64{12, 2, 3, 10, 5, 7, 8, 9, 9, 9, 3, 1}

	p1 := stats.Q1.Quartile(slice)
	p2 := stats.Q2.Quartile(slice)
	p3 := stats.Q3.Quartile(slice)

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

	p01 := stats.Percentile(slice, 0.01)
	p10 := stats.Percentile(slice, 0.1)
	p50 := stats.Percentile(slice, 0.5)
	p90 := stats.Percentile(slice, 0.9)
	p99 := stats.Percentile(slice, 0.99)

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

func ExampleP_Percentile() {
	slice := []float64{12, 2, 8, 3, 9, 7, 3, 9, 5, 9, 10, 1}

	r50 := stats.P50.Percentile(slice)
	r99 := stats.P99.Percentile(slice)

	fmt.Println(r50)
	fmt.Println(r99)
	// Output:
	// 8
	// 12
}
