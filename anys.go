// Package anys provides utilities for working with slices of any type.
package anys

// From returns a slice of type `any` containing the provided values.
//
// The From function is a utility that converts a slice of any type into a slice of `any`.
// This is particularly useful when you need to work with a homogeneous collection of elements
// of different types or when interfacing with APIs that require a slice of `any`.
//
// Example:
//
//	ii := []int{1, 2, 3}
//	aa := anys.From(ii...) // Converts the slice of integers into a slice of any
func From[S any](s ...S) []any {
	if s == nil {
		return nil
	}
	a := make([]any, len(s))
	for i, v := range s {
		a[i] = v
	}
	return a
}

// Map returns a slice of type T obtained by applying the conversion function
// to each element of the input values.
func Map[T, S any](conv func(S) T, s ...S) []T {
	if s == nil {
		return nil
	}
	a := make([]T, len(s))
	for i, v := range s {
		a[i] = conv(v)
	}
	return a
}

// BackTo returns a slice of type T by attempting to assert each element in the
// provided slice `s` to type T. If an element cannot be asserted to type T, it
// is set to the zero value of T.
//
// This function is useful for converting a slice of any (interface{}) values
// back to their original type.
//
// Example:
//
//	ii := []int{1, 2, 3}
//	aa := anys.From(ii...)
//	i2 := anys.BackTo[int](aa...) // Convert the slice of any back to a slice of int
func BackTo[T, S any](s ...S) []T {
	if s == nil {
		return nil
	}
	a := make([]T, len(s))
	for i, v := range s {
		if t, ok := any(v).(T); ok {
			a[i] = t
		}
		// If the assertion fails, a[i] remains the zero value of T, which is
		// already initialized by make.
	}
	return a
}

// Append takes a slice of any and appends the provided values to it.
//
// This function is particularly useful when you need to append elements from a
// slice of a specific type (e.g., []int) to a slice of type []any, which
// cannot be done directly using Go's built-in append function.
//
// Example usage:
//
//	a := []any{1, "two"}
//	ints := []int{3, 4, 5}
//	// The following line would cause a compile-time error because
//	// we cannot directly append a slice of int to a slice of any.
//	// a = append(a, ints...)
//	// Instead, use the custom Append function from the 'anys' package:
//	a = anys.Append(a, ints...)
//	fmt.Println(a)
//	// Output: [1 two 3 4 5]
func Append[S any](a []any, s ...S) []any {
	if s == nil {
		return a
	}
	a = grow(a, len(s))
	for _, v := range s {
		a = append(a, v)
	}
	return a
}

// AppendMap returns a slice of type T by applying the conversion function to
// each provided value and appending the result to the given slice.
func AppendMap[T, S any](t []T, conv func(S) T, s ...S) []T {
	if s == nil {
		return t
	}
	t = grow(t, len(s))
	for _, v := range s {
		t = append(t, conv(v))
	}
	return t
}

func grow[A any](a []A, n int) []A {
	if n -= cap(a) - len(a); n > 0 {
		a = append(a[:cap(a)], make([]A, n)...)[:len(a)]
	}
	return a
}

// AppendBackTo returns a slice of type T by appending values that have been type
// asserted to T from the provided arguments.
// If a value cannot be asserted to type T, the corresponding element is set to
// T's zero value.
func AppendBackTo[T, S any](t []T, s ...S) []T {
	if s == nil {
		return t
	}
	t = grow(t, len(s))
	var zero T
	for _, v := range s {
		if v, ok := any(v).(T); ok {
			t = append(t, v)
		} else {
			t = append(t, zero)
		}
	}
	return t
}
