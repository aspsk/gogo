package main

import (
	"fmt"
	"io"
	"os"
	"log"
)

type limitReader struct {
	r     io.Reader
	nread int64
	nmax  int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {

	nleft := lr.nmax - lr.nread
	if nleft == 0 {
		return 0, io.EOF
	}

	n = len(p)
	if int64(n) > nleft {
		n = int(nleft)
	}

	n, err = lr.r.Read(p[:n])
	if err != nil {
		return 0, err
	}

	lr.nread += int64(n)
	return n, nil

}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r: r, nmax: n}
}

func main() {

	buf := make([]byte, 1024)
	lrs := LimitReader(os.Stdin, 64)

	n, err := lrs.Read(buf)
	if err != nil {
		log.Fatalf("lrs.Read: %v", err)
	}
	fmt.Printf("read %d bytes:\n", n)
	fmt.Println(string(buf))
}
