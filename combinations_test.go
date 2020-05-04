package combinations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringCombinations(t *testing.T) {
	tt := []struct {
		name string
		in   []rune
		out  [][]rune
	}{
		{
			name: "Empty slice",
			in:   []rune{},
			out:  nil,
		},
		{
			name: "Single item",
			in:   []rune("A"),
			out: [][]rune{
				{'A'},
			},
		},
		{
			name: "Two items",
			in:   []rune("AB"),
			out: [][]rune{
				[]rune("A"),
				[]rune("B"),
				[]rune("AB"),
			},
		},
		{
			name: "Three items",
			in:   []rune("ABC"),
			out: [][]rune{
				[]rune("A"),
				[]rune("B"),
				[]rune("AB"),
				[]rune("C"),
				[]rune("AC"),
				[]rune("BC"),
				[]rune("ABC"),
			},
		},
		{
			name: "Four items",
			in:   []rune("ABCD"),
			out: [][]rune{
				[]rune("A"),
				[]rune("B"),
				[]rune("AB"),
				[]rune("C"),
				[]rune("AC"),
				[]rune("BC"),
				[]rune("ABC"),
				[]rune("D"),
				[]rune("AD"),
				[]rune("BD"),
				[]rune("ABD"),
				[]rune("CD"),
				[]rune("ACD"),
				[]rune("BCD"),
				[]rune("ABCD"),
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := All(tc.in)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}

func ExampleAll() {
	combinations := All([]string{"A", "B", "C"})
	fmt.Println(combinations)
	// Output:
	// [[A] [B] [A B] [C] [A C] [B C] [A B C]]
}

func TestStringCombinationsN(t *testing.T) {
	tt := []struct {
		name string
		in   []string
		n    int
		out  [][]string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			n:    1,
			out:  nil,
		},
		{
			name: "Single item",
			in:   []string{"A"},
			n:    1,
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "Two items, n = 0",
			in:   []string{"A", "B"},
			n:    0,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			name: "Two items, n = 1",
			in:   []string{"A", "B"},
			n:    1,
			out: [][]string{
				{"A"},
				{"B"},
			},
		}, {
			name: "Two items, n = 2",
			in:   []string{"A", "B"},
			n:    2,
			out: [][]string{
				{"A", "B"},
			},
		},
		{
			name: "Three items, n = 0",
			in:   []string{"A", "B", "C"},
			n:    0,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			name: "Three items, n = 1",
			in:   []string{"A", "B", "C"},
			n:    1,
			out: [][]string{
				{"A"},
				{"B"},
				{"C"},
			},
		},
		{
			name: "Three items, n = 2",
			in:   []string{"A", "B", "C"},
			n:    2,
			out: [][]string{
				{"A", "B"},
				{"A", "C"},
				{"B", "C"},
			},
		},
		{
			name: "Three items, n = 3",
			in:   []string{"A", "B", "C"},
			n:    3,
			out: [][]string{
				{"A", "B", "C"},
			},
		},
		{
			name: "Three items, n = 4",
			in:   []string{"A", "B", "C"},
			n:    4,
			out: [][]string{
				{"A", "B", "C"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := Combinations(tc.in, tc.n)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}
