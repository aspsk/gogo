package counters

import "io"
import "strings"

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	*c += WordCounter(len(strings.Fields(string(p))))
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	n := 0
	for _, x := range p {
		if x == '\n' {
			n++
		}
	}
	*c += LineCounter(n + 1)
	return len(p), nil
}

type WriterCounter struct {
	w io.Writer
	n int64
}

func (c *WriterCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	if err != nil {
		return n, err
	}
	c.n += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {

	ww := WriterCounter{w, 0}
	return &ww, &ww.n

}
