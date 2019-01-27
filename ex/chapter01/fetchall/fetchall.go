package main

import (
	"fmt"
	"io"
	"bufio"
	"io/ioutil"
	neturl "net/url"
	"net/http"
	"os"
	"strings"
	"time"
)

func uniqueFilename(name string) string {

	for i := 0; ; i++ {

		suffix := ""
		if i > 1 {
			suffix = fmt.Sprintf("(%d)", i)
		}

		if _, err := os.Stat(name + suffix); err == nil {
			continue
		}

		return name + suffix
	}

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	writer := ioutil.Discard
	u, err := neturl.Parse(url)
	if err == nil {
		os.Mkdir("out", 0755)
		filename := uniqueFilename("out/" + u.Host)
		f, err := os.Create(filename)
		if err == nil {
			defer f.Close()
			writer = bufio.NewWriter(f)
		} else {
			fmt.Printf("out/%s: %v [%v]\n", filename, err)
		}
	} else {
		fmt.Printf("%v\n", err)
	}

	nbytes, err := io.Copy(writer, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}

func main() {
	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", secs)
}
