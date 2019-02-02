# Exercise 7.4

The `strings.NewReader` function return a value that satisfies the `io.Reader`
interface (and others) by reading from its argument, a string. Implement a
simple version of `NewReader` yourself, and use it to make the HTML parser
(paragraph 5.2) take input from a string.
