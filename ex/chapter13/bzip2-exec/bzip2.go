package bzip

import (
	"bytes"
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	w    io.Writer
	buf  bytes.Buffer
	open bool
	sync.Mutex
}

func NewWriter(out io.Writer) io.WriteCloser {
	return &writer{w: out, buf: bytes.Buffer{}, open: true}
}

func (w *writer) Write(data []byte) (int, error) {
	w.Lock()
	defer w.Unlock()

	if !w.open {
		panic("closed")
	}

	w.buf.Write(data)
	return len(data), nil
}

func (w *writer) Close() error {
	w.Lock()
	defer w.Unlock()

	if !w.open {
		panic("closed")
	}
	defer func() { w.open = false }()

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("/usr/bin/bzip2", "-c", "-")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic("can't get stdin")
	}

	if err = cmd.Start(); err != nil {
		panic("can't start command")
	}

	io.Copy(stdin, &w.buf)
	stdin.Close()

	cmd.Wait()

	io.Copy(w.w, &stdout)

	return nil
}
