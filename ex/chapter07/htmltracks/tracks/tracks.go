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

var TrackList = []*Track{
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
	T []*Track
	history []string
	Maxdepth int
}

func (tl *Tracks) Len() int {
	return len(tl.T)
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
			if tl.T[i].Title != tl.T[j].Title {
				return b2i(tl.T[i].Title < tl.T[j].Title)
			}
		case "Artist":
			if tl.T[i].Artist != tl.T[j].Artist {
				return b2i(tl.T[i].Artist < tl.T[j].Artist)
			}
		case "Album":
			if tl.T[i].Album != tl.T[j].Album {
				return b2i(tl.T[i].Album < tl.T[j].Album)
			}
		case "Year":
			if tl.T[i].Year != tl.T[j].Year {
				return b2i(tl.T[i].Year < tl.T[j].Year)
			}
		case "Length":
			if tl.T[i].Length != tl.T[j].Length {
				return b2i(tl.T[i].Length < tl.T[j].Length)
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

	if len(tl.history) > tl.Maxdepth {
		tl.history = tl.history[:tl.Maxdepth]
	}
}

func (tl *Tracks) Swap(i, j int) {
	tl.T[i], tl.T[j] = tl.T[j], tl.T[i]
}
