package combinations

import (
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

// func ExampleAll() {
// 	combinations := All([]rune("ABC"))
// 	fmt.Println(combinations)
// 	// Output:
// 	// [[A] [B] [A B] [C] [A C] [B C] [A B C]]
// }

func TestStringCombinationsN(t *testing.T) {
	tt := []struct {
		name string
		in   []rune
		n    int
		out  [][]rune
	}{
		{
			name: "Empty slice",
			in:   []rune{},
			n:    1,
			out:  nil,
		},
		{
			name: "Single item",
			in:   []rune("A"),
			n:    1,
			out: [][]rune{
				[]rune("A"),
			},
		},
		{
			name: "Two items, n = 0",
			in:   []rune("AB"),
			n:    0,
			out: [][]rune{
				[]rune("A"),
				[]rune("B"),
				[]rune("AB"),
			},
		},
		{
			name: "Two items, n = 1",
			in:   []rune("AB"),
			n:    1,
			out: [][]rune{
				[]rune("A"),
				[]rune("B"),
			},
		}, {
			name: "Two items, n = 2",
			in:   []rune("AB"),
			n:    2,
			out: [][]rune{
				[]rune("AB"),
			},
		},
		{
			name: "Three items, n = 0",
			in:   []rune("ABC"),
			n:    0,
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
			name: "Three items, n = 1",
			in:   []rune("ABC"),
			n:    1,
			out: [][]rune{
				[]rune("A"),
				[]rune("B"),
				[]rune("C"),
			},
		},
		{
			name: "Three items, n = 2",
			in:   []rune("ABC"),
			n:    2,
			out: [][]rune{
				[]rune("AB"),
				[]rune("AC"),
				[]rune("BC"),
			},
		},
		{
			name: "Three items, n = 3",
			in:   []rune("ABC"),
			n:    3,
			out: [][]rune{
				[]rune("ABC"),
			},
		},
		{
			name: "Three items, n = 4",
			in:   []rune("ABC"),
			n:    4,
			out: [][]rune{
				[]rune("ABC"),
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
