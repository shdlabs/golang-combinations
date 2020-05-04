package combinations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleAll() {
	combinations := All("ABC")
	fmt.Println(combinations)
	// Output:
	// [A B AB C AC BC ABC]
}
func TestStringCombinations(t *testing.T) {
	tt := []struct {
		name string
		in   string
		out  []string
	}{
		{
			name: "Empty slice",
			in:   "",
			out:  nil,
		},
		{
			name: "Single item",
			in:   "A",
			out:  []string{"A"},
		},
		{
			name: "Two items",
			in:   "AB",
			out:  []string{"A", "B", "AB"},
		},
		{
			name: "Three items",
			in:   "ABC",
			out:  []string{"A", "B", "AB", "C", "AC", "BC", "ABC"},
		},
		{
			name: "Four items",
			in:   "ABCD",
			out: []string{"A", "B", "AB", "C", "AC", "BC",
				"ABC", "D", "AD", "BD", "ABD", "CD", "ACD", "BCD", "ABCD"},
		},
	}
	assert := assert.New(t)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := All(tc.in)
			assert.Equal(out, tc.out)
		})
	}
}

func TestStringCombinationsN(t *testing.T) {
	tt := []struct {
		name string
		in   string
		n    int
		out  []string
	}{
		{
			name: "Empty slice",
			in:   "",
			n:    1,
			out:  nil,
		},
		{
			name: "Single item",
			in:   "A",
			n:    1,
			out:  []string{"A"},
		},
		{
			name: "Two items, n = 0",
			in:   "AB",
			n:    0,
			out:  []string{"A", "B", "AB"},
		},
		{
			name: "Two items, n = 1",
			in:   "AB",
			n:    1,
			out:  []string{"A", "B"},
		}, {
			name: "Two items, n = 2",
			in:   "AB",
			n:    2,
			out:  []string{"AB"},
		},
		{
			name: "Three items, n = 0",
			in:   "ABC",
			n:    0,
			out:  []string{"A", "B", "AB", "C", "AC", "BC", "ABC"},
		},
		{
			name: "Three items, n = 1",
			in:   "ABC",
			n:    1,
			out:  []string{"A", "B", "C"},
		},
		{
			name: "Three items, n = 2",
			in:   "ABC",
			n:    2,
			out:  []string{"AB", "AC", "BC"},
		},
		{
			name: "Three items, n = 3",
			in:   "ABC",
			n:    3,
			out:  []string{"ABC"},
		},
		{
			name: "Three items, n = 4",
			in:   "ABC",
			n:    4,
			out:  []string{"ABC"},
		},
	}
	assert := assert.New(t)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := Combinations(tc.in, tc.n)
			assert.Equal(out, tc.out)
		})
	}
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		All("ABCD")
	}
}

func BenchmarkCombinations(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Combinations("ABCD", 2)
	}
}
