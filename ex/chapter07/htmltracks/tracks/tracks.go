package tracks

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type Tracks struct {
	t []*Track
	history []string
	maxdepth int
}

func (tl *Tracks) Len() int {
	return len(tl.t)
}

func b2i(b bool) int {
	if b {
		return -1
	} else {
		return 1
	}
 }

func (tl *Tracks) less(column string, i, j int) int {
	switch column {
		case "Title":
			if tl.t[i].Title != tl.t[j].Title {
				return b2i(tl.t[i].Title < tl.t[j].Title)
			}
		case "Artist":
			if tl.t[i].Artist != tl.t[j].Artist {
				return b2i(tl.t[i].Artist < tl.t[j].Artist)
			}
		case "Album":
			if tl.t[i].Album != tl.t[j].Album {
				return b2i(tl.t[i].Album < tl.t[j].Album)
			}
		case "Year":
			if tl.t[i].Year != tl.t[j].Year {
				return b2i(tl.t[i].Year < tl.t[j].Year)
			}
		case "Length":
			if tl.t[i].Length != tl.t[j].Length {
				return b2i(tl.t[i].Length < tl.t[j].Length)
			}
	}
	return 0
}

func (tl *Tracks) Less(i, j int) bool {

	for _, column := range tl.history {
		x := tl.less(column, i, j)
		if x == 0 {
			continue
		}
		return x < 0
	}
	return false
}

func (tl *Tracks) AddHistory(column string) {
	if tl.history == nil || len(tl.history) == 0 {
		tl.history = []string{"Title"}
	}

	if column != tl.history[0] {
		n := make([]string, len(tl.history)+1)
		n[0] = column
		copy(n[1:], tl.history)
		tl.history = n
	}

	if len(tl.history) > tl.maxdepth {
		tl.history = tl.history[:tl.maxdepth]
	}
}

func (tl *Tracks) Swap(i, j int) {
	tl.t[i], tl.t[j] = tl.t[j], tl.t[i]
}
