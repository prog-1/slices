# Exercises: Slices

> All exercises should contain a runnable application and tests!

> Every program must contain tests and must be created in a dedicated folder.
Program names are provided in the task description.

1. Check if a slice contains at least a single number which is:
    - odd (`AnyIsOdd(s []int) bool`);
    - even (`AnyIsEven(s []int) bool`).

   Program name: `any_odds_and_evens`.

2. Check if a slice contains only numbers which are:
    - odd (`AreOddsOnly(s []int) bool`);
    - even (`AreEvenOnly(s []int) bool`).

   Program name: `only_odds_and_evens`.

3. Create a copy of a slice, containing only numbers which are:
    - odd (`FilterOdds(s []int) []int`);
    - even (`FilterEvens(s []int) []int`).

   Program name: `filter_odds_and_evens`.

4. Sort a slice using:
    - a custom bubble-sort (or any other algorithm) (`Sort1(s []int)`);
    - built-in `sort.Ints([]int)` (`Sort2(s []int)`).

   Program name: `sort`.

5. Create a two-dimensional 10x10 array which contains a multiplication table
   (`MultTable() [10][10]int`).

   Program name: `mult_table`.

6. Given width and height construct a two-dimensional slice representing a
   rectangle (`Rect(x, y int) [][]int`). Fill its borders with `1`. For width=7
   and height=5 it would be:

   ```
   1111111
   1000001
   1000001
   1000001
   1111111
   ```

   Program name: `rect`.

7. Given width and height construct a two-dimensional slice representing a
   rectangle (`Spiral(x, y int) [][]int`). Fill it with incremental numbers
   starting from 1. For width=7 and height=5 it should look like this:

   ```
    1  2  3  4  5  6  7
   20 21 22 23 24 25  8
   19 32 33 34 35 26  9
   18 31 30 29 28 27 10
   17 16 15 14 13 12 11
   ```

   Program name: `spiral`.