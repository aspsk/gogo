# Exercise 7.8

Many GUIs provide a table widget with a stateful multi-tier sort: the primary
sort key is the most recently clicked column head, the secondary sort key is
the second-most recently clicked column head, and so on. Define an
implementation of `sort.Interface` for use by such a table. Compare that
approach with repeated sorting using `sort.Stable`.

# Exercise 7.9

Use the `html/template` package (paragraph 4.6) to replace `printTracks` with a
function that displays the tracks as an HTML table. Use the solution to the
previous exercise to arrange that each click on a column head makes an HTTP
request to sort the table.
