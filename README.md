# anys

A Go package that provides utilities for working with slices of any type,
simplifying type conversions between concrete types and `any` (interface{}).

## Installation

```bash
go get github.com/goaux/anys
```

## Overview

The `anys` package offers a collection of functions that make it easier to work
with slices of various types, particularly when converting between concrete
types and the `any` type. This is useful when interfacing with APIs that
require a slice of `any` or when working with heterogeneous collections.

## Features

- Convert slices of specific types to slices of `any`
- Convert slices of `any` back to slices of specific types
- Apply conversion functions across slices
- Append values from one slice type to another
- Safely handle type conversions with zero value fallbacks

## Usage

### Converting a slice to `any`

```go
import (
	"fmt"
	"github.com/goaux/anys"
)

func main() {
	ints := []int{1, 2, 3}

	// Convert a slice of int to a slice of any
	anySlice := anys.From(ints...)

	fmt.Println(anySlice) // Output: [1 2 3]
}
```

### Converting a slice of `any` back to a specific type

```go
import (
	"fmt"
	"github.com/goaux/anys"
)

func main() {
	anySlice := []any{1, 2, 3}

	// Convert a slice of any back to a slice of int
	intSlice := anys.BackTo[int](anySlice...)

	fmt.Println(intSlice) // Output: [1 2 3]
}
```

### Mapping a function across a slice

```go
import (
	"fmt"
	"strconv"
	"strings"
	"github.com/goaux/anys"
)

func main() {
	ints := []int{1, 2, 3}

	// Map each integer to its string representation
	strSlice := anys.Map(func(i int) string {
		return strconv.Itoa(i)
	}, ints...)

	result := strings.Join(strSlice, ",")
	fmt.Println(result) // Output: 1,2,3
}
```

### Appending values of different types

```go
import (
	"fmt"
	"github.com/goaux/anys"
)

func main() {
	a := []any{1, "two"}
	ints := []int{3, 4, 5}

	// Append a slice of int to a slice of any.
	// The following line would cause a compile-time error because we cannot
	// directly append a slice of int to a slice of any.
	// a = append(a, ints...)
	a = anys.Append(a, ints...)

	fmt.Println(a) // Output: [1 two 3 4 5]
}
```

### Appending with type assertion

```go
import (
	"fmt"
	"github.com/goaux/anys"
)

func main() {
	ints := []int{1, 2}
	values := []any{3, 4, "invalid"}

	// Append values that can be type-asserted to int
	result := anys.AppendBackTo[int](ints, values...)

	fmt.Println(result) // Output: [1 2 3 4 0]
	// Note: "invalid" cannot be asserted to int, so it becomes 0 (zero value)
}
```

### Appending with conversion

```go
import (
	"fmt"
	"github.com/goaux/anys"
)

func main() {
	strings := []string{"a", "b"}
	ints := []int{1, 2}

	// Append ints to strings with conversion
	result := anys.AppendMap(strings, func(i int) string {
		return fmt.Sprintf("x%d", i)
	}, ints...)

	fmt.Println(result) // Output: [a b x1 x2]
}
```
