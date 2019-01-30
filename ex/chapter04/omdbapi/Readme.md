# Exercise 4.13

The JSON-based web service of the Open Movie Database lets you search
`https://omdbapi.com/` for a movie by name and download its poster image. Write
a tool `poster` that downloads the poster image for the movie named on the
command line.

# Solution

Now they disabled free access to posters, so I implemented just a simple
request to their API. One needs to register and obtain an access key to use it,
like this:
```
$ curl -s http://www.omdbapi.com/?apikey=xxxxxxxx\&t=terminator | jq
{
  "Title": "Terminator",
  "Year": "1991",
  "Rated": "N/A",
  "Released": "N/A",
  "Runtime": "39 min",
  "Genre": "Short, Action, Sci-Fi",
  "Director": "Ben Hernandez",
  "Writer": "James Cameron (characters), James Cameron (concept), Ben Hernandez (screenplay)",
  "Actors": "Loris Basso, James Callahan, Debbie Medows, Michelle Kovach",
  ...
```
