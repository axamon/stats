package stats

import "fmt"

func ExampleQ_Quartile() {

	slice := []float64{12, 2, 3, 10, 5, 7, 8, 9, 9, 9, 3, 1}

	p1 := q1.Quartile(slice)
	p2 := q2.Quartile(slice)
	p3 := q3.Quartile(slice)

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

	i := InterQuartileRange(slice)

	fmt.Println(i)
	// Output:
	// 6
}

func ExamplePercentile() {
	slice := []float64{1, 2, 3, 3, 5, 7, 8, 9, 9, 9, 10, 12}

	p01 := Percentile(slice, 0.01)
	p10 := Percentile(slice, 0.1)
	p50 := Percentile(slice, 0.5)
	p90 := Percentile(slice, 0.9)
	p99 := Percentile(slice, 0.99)

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
	slice := []float64{1, 2, 3, 3, 5, 7, 8, 9, 9, 9, 10, 12}

	r50 := p50.Percentile(slice)
	r99 := p99.Percentile(slice)

	fmt.Println(r50)
	fmt.Println(r99)
	// Output:
	// 8
	// 12
}
