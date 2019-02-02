package main

import (

	"fmt"
	"net/http"
	"log"
	"sort"
	"html/template"
	"github.com/aspsk/gogo/ex/chapter07/htmltracks/tracks"

)

var tpl = template.Must(template.New("issuelist").Parse(`
<h1>Tracks</h1>
<table>
<tr style='text-align: left'>
  <th><a href='.?column=Title'>Title</a></th>
  <th><a href='.?column=Artist'>Artist</a></th>
  <th><a href='.?column=Album'>Album</a></th>
  <th><a href='.?column=Year'>Year</a></th>
  <th><a href='.?column=Length'>Length</a></th>
</tr>
{{range .T}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

var tl *tracks.Tracks = &tracks.Tracks{T: tracks.TrackList, Maxdepth: 2}

func handler(w http.ResponseWriter, r *http.Request) {
	if x := r.FormValue("column"); x != "" {
		tl.AddHistory(x)
		fmt.Printf("column=%s\n", x)
	}

	sort.Sort(tl)
	if err := tpl.Execute(w, tl); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
