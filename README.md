# maduce

maduce is an implementation of the functional concepts filter, map
and reduce found in other languages like python, javascript, etc.

This package purposely diverge from core principals of how go code should be
written, you should therefore think twice before you consider using this
package, in most cases for loops is the way to go.

The reason for the existence of this package is that it allows for better
composability and allows datasets to be more easily explored and evaluated in
Go. It's specifically designed as a tool to be used for experimenting with
datasets and not as a library intented for production use where performance
is critical.

The API of the maduce package is completely dynamic which has the down side
of no compile time garuantees about the function signatures given to filter,
map or reduce.

Each method on a `maduce.Collection` have a description of the handlers they
support.
Because go doesn't support generics yet, i have create my own notation where
`<Type>` can be replaced with what ever type you want. The `<Type>` in the
function argument has to be the same as in the collection. The output type
could be something else or the same as the input, it depends on what you want
to achieve.
```go
// example of function signature with generic types
func(item <Type>, index int) <Type>

// example of function that maps over a collection of float64 and castst them
// to a string in a new collection, this example satisfies the function
// signature from above
func(item float64, index int) string {
	return fmt.Sprintf("%.2f", item)
}
```
This package is heavily based on reflection and type assertions which can
result in runtime panics if used wrongly.

`TODO(@kvartborg)`: would like to experiment with a streaming implementation
based on the io.Reader interface at some point.

## Usage
Below is a simple example of how a slice of `float64` can be filtered, mapped
and reduced into a string.
```go
package main

import (
    "fmt"

    "github.com/kvartborg/maduce"
)

func main() {
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
}
```

## Documentation
The full documentation of the package can be found on [godoc](https://pkg.go.dev/github.com/kvartborg/maduce?tab=doc).

## License
This project is licensed under the [MIT License](https://github.com/kvartborg/maduce/blob/master/LICENSE).
