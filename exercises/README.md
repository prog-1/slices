# Exercises: Slices

IMPORTANT: All exercises should contain a runnable application and tests!

1. Check if a slice contains at least a single number which is:
    - odd;
    - even.  
2. Check if a slice contains only numbers which are:
    - odd;
    - even.
3. Create a copy of a slice, containing only numbers which are:
    - odd;
    - even.
4. Sort a slice using:
    - a custom bubble-sort (or any other algorithm);
    - built-in `sort.Ints([]int)`.
5. Create a two-dimensional 10x10 array which contains a multiplication table.
6. Given width and height construct a two-dimensional slice representing a
   rectangle. Fill its borders with `1`. For width=7 and height=5 it would be:
   ```
   1111111
   1000001
   1000001
   1000001
   1111111
   ```
7. Given width and height construct a two-dimensional slice representing a
   rectangle. Fill it with incremental numbers starting from 1. For width=7 and
   height=5 it should look like this:
   ```
    1  2  3  4  5  6  7
   20 21 22 23 24 25  8
   19 32 33 34 35 26  9
   18 31 30 29 28 27 10
   17 16 15 14 13 12 11
   ```