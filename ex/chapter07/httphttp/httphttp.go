package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type database struct {
	items map[string]int
	sync.Mutex
}

var tpl = template.Must(template.New("issuelist").Parse(`
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $k, $v := .}}
<tr>
  <td>{{$k}}</td>
  <td>{{$v}}</td>
</tr>
{{end}}
</table>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()
	if err := tpl.Execute(w, db.items); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()
	item := req.URL.Query().Get("item")
	if price, ok := db.items[item]; ok {
		fmt.Fprintf(w, "$%d\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db.items[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	price := req.URL.Query().Get("price")
	x, err := strconv.Atoi(price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad price parameter: %q\n", price)
		return
	}

	db.items[item] = x
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db.items[item]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item doesn't exist: %q\n", item)
		return
	}

	price := req.URL.Query().Get("price")
	x, err := strconv.Atoi(price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad price parameter: %q\n", price)
		return
	}

	db.items[item] = x
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db.items[item]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item doesn't exist: %q\n", item)
		return
	}

	delete(db.items, item)
}

func main() {
	db := database{items: map[string]int{"shoes": 50, "socks": 5},}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
