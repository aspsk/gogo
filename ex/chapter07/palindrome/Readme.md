# Exercise 7.10

The `sort.Interface` type can be adapted to other uses. Write a function
`IsPalindrome(s sort.Interface) bool` that reports whether the sequence `s`
is a palindrome. Assume that elements at indices `i` and `j` are equal if
`!s.Less(i, j) && !s.Less(j, i)`

# Exercise 11.3

`TestRandomPalindromes` only tests palindromes. Write a randomized test that
generates and verifies _non_-palindromes.

# Exercise 11.4

Modify `randomPalindrome to exercise `IsPalindrome`'s handling of punctuation
and spaces.
