# Arrays & Slices

https://go.dev/blog/slices-intro

## Arrays

The type `[n]T` in an array of `n` values of type `T` in Go. Arrays cannot be
resized in Go. Each value in the array could be addressed using `[n]` notation,
arrays elements are allocated in a linear memory, and it is very similar to
defining variables of the same type, but we address those different variables
through a single name and an index.

Consider this example:

  ```go
  var a0, a1, a2, a3 int
  a1 = 123
  fmt.Println(a0, a1, a2, a3) // Output: 0 123 0 0
  ```

This could be rewritten using arrays:

  ```go
  var a [4]int
  a[1] = 123
  fmt.Println(a) // Output: [0 123 0 0]
  ```

## Slices

Slices is an abstraction on top of arrays. The abstraction provides a "view" of
elements in a linear subset of elements of an array in the range `[start; end)`.
E.g. `s := a[1:3]` from the last example would refer elements with indices 1 and
2.

Demonstration of an array `a` and a slice of it `s`:

```go
a := [4]int{1, 2, 3, 4}
s := a[1:3]
fmt.Println(a) // Output: [1 2 3 4]
fmt.Println(s) // Output: [2 3]

a[2] = 0
fmt.Println(a) // Output: [1 2 0 4]
fmt.Println(s) // Output: [2 0]

s[0] = -1
fmt.Println(a) // Output: [1 -1 0 4]
fmt.Println(s) // Output: [-1 0]
```

Slices are more convenient than arrays, since their length can be changed. It is
also important to understand, that slices are built on top of arrays, so if we
want to add a new element to a slice, and underlying array has enough capacity,
no new memory will be allocated. Otherwise, Go would allocate a new array, copy
existing elements and return a new slice.

`len([]T)` returns the length of an array of a slice. `cap([]T)` returns the
capacity. For arrays these two are always the same. For slices they both can change.

`[]T append([]T, v T, [...])` appends new values to slices. Check this example:

```go
fmt.Println(len(a), cap(a)) // Output: 4 4
fmt.Println(len(s), cap(s)) // Output: 2 3

s = append(s, -1)
fmt.Println(len(a), cap(a), a) // Output: 4 4 4 4 [1 -1 0 -1]
fmt.Println(len(s), cap(s), s) // Output: 3 3 [-1 0 -1]

s = append(s, 123)
fmt.Println(len(a), cap(a), a) // Output: 4 4 [1 -1 0 -1]
fmt.Println(len(s), cap(s), s) // Output: 4 6 [-1 0 -1 123]
```

## For-range

It is possible to iterate over elements using indices and any of the `for` loops
that we know. E.g.

```go
// AnyIsOdd returns true iff any number in the slice is odd.
func AnyIsOdd(s []int) bool {
  for i := 0; i < len(s); i++ {
    if s[i] % 2 == 1 {
      return true
    }
  }
  return false
}

func main() {
  fmt.Println(AnyIsOdd([]int{2, 4, 0})) // Output: false
  fmt.Println(AnyIsOdd([]int{2, 1, 0})) // Output: true
}
```

However, it is very common to iterate over all of the elements one-by-one. It is
so common, that Go has a dedicated looping construction called `for-range`:

```go
// AnyIsEven returns an index of the first odd number in the collection, or -1
// if there are no even numbers.
func IndexOfEven(s []int) int {
  for i, v := range s {
    if v % 2 == 0 {
      return i
    }
  }
  return -1
}

func main() {
  fmt.Println(IndexOfEven([]int{1, 17, 3}))     // Output: -1
  fmt.Println(IndexOfEven([]int{1, 17, 6, 98})) // Output: 2
}
```

## Exercises

Check the [exercises](exercises) folder for the list of problems.