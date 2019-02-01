# Exercise 6.1

Implement these additional methods:
```
    func (*IntSet) Len() int      // return the number of elements
    func (*IntSet) Remove(x int)  // remove x from the set
    func (*IntSet) Clear()        // remove all elements from the set
    func (*IntSet) Copy() *IntSet // return a copy of the set
```

# Exercise 6.2

Define a variadic `(*IntSet).AddAll(...int)` method that allows a list of
values to be added, such as `s.AddAll(1, 2, 3)`.

# Exercise 6.3

`(*IntSet).UnionWith` computes the union of two sets using `|`, the
world-parallel bitwise OR operator. Implement methods for `IntersectWith`,
`DifferenceWith`, and `SymmetricDifference` for the corresponding set
operations.

# Exercise 6.4

Add a method `Elems` that returns a slice containing the elements of the set,
suitable for iterating over with a range loop.

# Exercise 6.5

The type of each word used by `IntSet` is `uint64`, but 64-bit arithmetic may
be inefficient on a 32-bit platform. Modify the program to use the `uint` type,
which is the most efficient unsigned integer type for the platform.
