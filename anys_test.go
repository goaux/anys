package anys_test

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/goaux/anys"
)

func ExampleFrom() {
	ints := []int{3, 4, 5}
	x := anys.From(ints...) // From is equivalent to the special case of Map
	y := anys.Map(func(i int) any { return any(i) }, ints...)
	fmt.Println(reflect.DeepEqual(x, y))
	// Output:
	// true
}

func TestFrom(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []any
	}{
		{[]int{1, 2, 3}, []any{1, 2, 3}},
		{[]int{}, []any{}},
		{[]int{1, 2, 3, 4, 5}, []any{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		actual := anys.From(tc.input...)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("From(%v) = %v, expected %v", tc.input, actual, tc.expected)
		}
	}

	t.Run("nil", func(t *testing.T) {
		actual := anys.From[int]()
		if actual != nil {
			t.Errorf("From[int]() = %v, expected nil", actual)
		}
	})
}

func ExampleMap() {
	ii := []int{1, 2, 3}
	s := strings.Join(
		anys.Map(func(i int) string { return strconv.Itoa(i) }, ii...),
		",",
	)
	fmt.Println(s)
	// Output:
	// 1,2,3
}

func TestMap(t *testing.T) {
	testCases := []struct {
		input    []int
		conv     func(int) string
		expected []string
	}{
		{[]int{1, 2, 3}, func(i int) string { return fmt.Sprintf("%d", i) }, []string{"1", "2", "3"}},
		{[]int{}, func(i int) string { return fmt.Sprintf("%d", i) }, []string{}},
		{[]int{1, 2, 3, 4, 5}, func(i int) string { return fmt.Sprintf("x%d", i) }, []string{"x1", "x2", "x3", "x4", "x5"}},
	}

	for _, tc := range testCases {
		actual := anys.Map(tc.conv, tc.input...)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Map(%v, conv) = %v, expected %v", tc.input, actual, tc.expected)
		}
	}

	t.Run("nil", func(t *testing.T) {
		actual := anys.Map(func(int) rune { return ' ' })
		if actual != nil {
			t.Errorf("Map() = %v, expected nil", actual)
		}
	})
}

func ExampleBackTo() {
	ii := []int{1, 2, 3}
	aa := anys.From(ii...)
	i2 := anys.BackTo[int](aa...)
	fmt.Println(reflect.DeepEqual(ii, i2))
	// Output:
	// true
}

func TestBackTo(t *testing.T) {
	testCases := []struct {
		input    []any
		expected []int
	}{
		{[]any{1, 2, 3}, []int{1, 2, 3}},
		{[]any{}, []int{}},
		{[]any{1, "2", 3.0, 4}, []int{1, 0, 0, 4}},
		{[]any{1.5, 2.7, 3.9}, []int{0, 0, 0}},
	}

	for _, tc := range testCases {
		actual := anys.BackTo[int](tc.input...)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("BackTo(%v) = %v, expected %v", tc.input, actual, tc.expected)
		}
	}

	t.Run("nil", func(t *testing.T) {
		actual := anys.BackTo[int, any]()
		if actual != nil {
			t.Errorf("BackTo[int]() = %v, expected nil", actual)
		}
	})
}
func ExampleAppend() {
	a := []any{1, "two"}
	ints := []int{3, 4, 5}
	// The following line would cause a compile-time error because we cannot
	// directly append a slice of int to a slice of any.
	// a = append(a, ints...)
	// Instead, use the custom Append function from the 'anys' package to achieve this:
	a = anys.Append(a, ints...)
	fmt.Println(a)
	// Output: [1 two 3 4 5]
}

func TestAppend(t *testing.T) {
	testCases := []struct {
		initial  []any
		add      []int
		expected []any
	}{
		{[]any{1, 2}, []int{3, 4}, []any{1, 2, 3, 4}},
		{[]any{}, []int{1, 2, 3}, []any{1, 2, 3}},
		{[]any{1, 2, 3}, []int{}, []any{1, 2, 3}},
		{[]any{"a", "b"}, []int{1, 2}, []any{"a", "b", 1, 2}},
	}

	for _, tc := range testCases {
		actual := anys.Append(tc.initial, tc.add...)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Append(%v, %v) = %v, expected %v", tc.initial, tc.add, actual, tc.expected)
		}
	}

	t.Run("nil", func(t *testing.T) {
		initial := []any{1, 2}
		actual := anys.Append[int](initial)
		expected := initial
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Append(%v) = %v, expected %v", initial, actual, expected)
		}
	})
}

func TestAppendMap(t *testing.T) {
	testCases := []struct {
		initial  []string
		add      []int
		conv     func(int) string
		expected []string
	}{
		{[]string{"a", "b"}, []int{1, 2}, func(i int) string { return fmt.Sprintf("x%d", i) }, []string{"a", "b", "x1", "x2"}},
		{[]string{}, []int{1, 2, 3}, func(i int) string { return fmt.Sprintf("x%d", i) }, []string{"x1", "x2", "x3"}},
		{[]string{"a", "b", "c"}, []int{}, func(i int) string { return fmt.Sprintf("x%d", i) }, []string{"a", "b", "c"}},
	}

	for _, tc := range testCases {
		actual := anys.AppendMap(tc.initial, tc.conv, tc.add...)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("AppendMap(%v, conv, %v) = %v, expected %v", tc.initial, tc.add, actual, tc.expected)
		}
	}

	t.Run("nil", func(t *testing.T) {
		type Int int
		initial := []int{1, 2}
		actual := anys.AppendMap(initial, func(string) int { return 0 })
		expected := initial
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("AppendMap(%v) = %v, expected %v", initial, actual, expected)
		}
	})
}

func TestAppendBackTo(t *testing.T) {
	testCases := []struct {
		initial  []int
		add      []any
		expected []int
	}{
		{[]int{1, 2}, []any{3, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []any{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []any{}, []int{1, 2, 3}},
		{[]int{1, 2}, []any{"a", "b"}, []int{1, 2, 0, 0}},
	}

	for _, tc := range testCases {
		actual := anys.AppendBackTo(tc.initial, tc.add...)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("AppendBackTo(%v, %v) = %v, expected %v", tc.initial, tc.add, actual, tc.expected)
		}
	}

	t.Run("nil", func(t *testing.T) {
		type Int int
		initial := []int{1, 2}
		actual := anys.AppendBackTo[int, Int](initial)
		expected := initial
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("AppendBackTo(%v) = %v, expected %v", initial, actual, expected)
		}
	})
}
