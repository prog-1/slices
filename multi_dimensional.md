# Multi-Dimensional Arrays and Slices

https://golangbyexample.com/two-dimensional-array-slice-golang/

## Multi-Dimensional Arrays

Multi-demensional array can be declared in Go as follows:

```go
[len1][len2][len3]....[lenN]T{}
```

where

 * `len1`, `len2` .. `lenN` are length of each of the dimensions
 * `T` is the data type.

### Examples

Declaring an **empty** 2d array with 4 rows and 5 columns:

```go
var table [4][5]int
//  [
//      [0, 0, 0, 0, 0],
//      [0, 0, 0, 0, 0],
//      [0, 0, 0, 0, 0],
//      [0, 0, 0, 0, 0],
//  ]
```

Declaring a **non-empty** 2d array with 4 rows and 5 columns:

```go
table := [4][5]int{
    {1, 2, 3, 4, 5},
    {6, 7, 8, 9, 10},
    {11, 12, 13, 14, 15},
    {16, 17, 18, 19, 20},
}
```

Accessing the number of rows, columns and elements:

```go
fmt.Printf("Number of rows: %d\n", len(table)) // 4
fmt.Printf("Number of columns: %d\n", len(table[0])) // 5
fmt.Printf("Total number of elements: %d\n", len(table)*len(table[0])) // 20
```

Modifying and accessing array elements:

```go
table[0][2] = -1
fmt.Println(table[0][2]) // -1
fmt.Println(table[0])    // [1 2 -1 4 5]

for i := range table {
    for j := range table[i] {
        table[i][j]++
    }
}
// [
//  [2 3 0 5 6],
//  [7 8 9 10 11],
//  [12 13 14 15 16],
//  [17 18 19 20 21],
// ]
```

We have to use a nested range for traversal using a for-range loop. The first
range traverses each of the rows. The second range traverses the individual
array present at that row.

Pretty printing:

```go
for row := 0; row < 4; row++ {
    for col := 0; col < 5; col++ {
        fmt.Printf("%3d", table[row][col])
    }
    fmt.Println()
}
//  1  2  0  4  5
//  6  7  8  9 10
// 11 12 13 14 15
// 16 17 18 19 20
```

## Multi-Dimensional Slices

In contast to the array declaration, a slice declaration **does not contain
dimension lengths**.

```go
[][][]...[]T{}
```

where `T` is the data type.

### Examples

Memory for slice elements has to be initialized explicitly using `make([][][]...[]T{}, len1)`, where `len1` is the length of the outermost dimension.

```go
var [][]int s
fmt.Println(s) // []

rows, cols := 4, 5
s := make([][]int, rows)
fmt.Println(s) // [[] [] [] []]

table[0] = []int{1, 2, 3, 4, 5}
for row := 1; row < rows; row++ {
    table[row] = make([]int, cols)
}
//  [
//      [1, 2, 3, 4, 5],
//      [0, 0, 0, 0, 0],
//      [0, 0, 0, 0, 0],
//      [0, 0, 0, 0, 0],
//  ]
```