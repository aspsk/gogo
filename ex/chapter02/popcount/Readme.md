# Exercise 2.3

Rewrite `PopCount` to use a loop instead of a single expression. Compare the
performance of the two versions.

# Exercise 2.4

Write a version of `PopCount` that counts bits by shifting its argument through
64 bit positions, testing the rightmost bit each time. Compare its performance
to the tables-lookup version.

# Exercise 2.5

The expression `x & (x-1)` clears the rightmost non-zero bit of x. Write a
version of `PopCount` that counts bits using this fact, and assess its
performance.

# Solution

Wrote the following functions:

  * `PopCount`: the original from the book
  * `PopCount2`: the same as `PopCount`, but uses a table of size 65536
  * `PopCountLoop`: uses a loop (exercise 2.3)
  * `PopCountShift`: uses a shift (exercise 2.4)
  * `PopCountRightmost`: uses the `x&(x-1)` expression (exercise 2.5)
  * `PopCountC`: a C function which uses the `POPCNT` HW instruction

## Benchmarks

I wanted to compare all the implementations with the native `POPCNT` hardware
operation which is available in C as `__builtin_popcntll` (for 64-bit values).
So I wrote a function, `PopCountC` which had shown the worse results (68ns).
Then I wrote a function which executes the operation 10,000,000 times so that
we don't count expenses of executing C code from Go, and made all other
benchmarks to execute 10,000,000 times as well. Now the results are more
meaningful (sorted from the best to worse):

|Function|ns/10,000,000 operations|ns/operation|
| --- | ---: | ---: |
| `PopCountC10000000` |  18414958 |  1.8 |
| `PopCount2`         |  28863041 |  2.9 |
| `PopCount`          |  45487204 |  4.5 |
| `PopCountRightmost` |  47891559 |  4.8 |
| `PopCountLoop`      | 192832530 | 19.3 |
| `PopCountShift`     | 442337802 | 44.2 |
| `PopCountC`         | 680730800 | 68.1 |

# Exercise 11.6

Write benchmarks to compare the `PopCount` implementation in Section 2.6.2 with
your solutions to Exercise 2.4 and Exercise 2.5.
