package maduce_test

import (
	"fmt"
	"strconv"

	"github.com/kvartborg/maduce"
)

type Product struct {
	ID    int
	Title string
	Price float64
}

func Example() {
	collection := maduce.From([]float64{0, 1.8, 2, 3.3, 4, 5})

	var result string
	collection.
		Filter(func(n float64) bool {
			return n > 0
		}).
		Map(func(n float64) string {
			return fmt.Sprintf("%.2f", n)
		}).
		Reduce(&result, func(s, result string, index int) string {
			if result == "" {
				return fmt.Sprintf("%d: %s", index, s)
			}
			return fmt.Sprintf("%s\n%d: %s", result, index, s)
		})

	fmt.Println(result)
	// Output:
	// 0: 1.80
	// 1: 2.00
	// 2: 3.30
	// 3: 4.00
	// 4: 5.00
}

func ExampleCollection_Filter() {
	collection := maduce.From([]int{1, 2, 3, 4, 5})

	result := collection.Filter(func(n int) bool {
		return n > 3
	})

	fmt.Println(result)
	// Output: [4 5]
}

func ExampleCollection_Map() {
	collection := maduce.From([]int{1, 2, 3, 4, 5})

	result := collection.Map(func(n int) string {
		return strconv.Itoa(n)
	})

	fmt.Printf("%#v\n", result)
	// Output: maduce.Collection{"1", "2", "3", "4", "5"}
}

func ExampleCollection_Reduce() {
	collection := maduce.From([]int{1, 2, 3, 4, 5})

	var sum int
	collection.Reduce(&sum, func(n, sum int) int {
		return n + sum
	})

	fmt.Println(sum)
	// Output: 15
}
