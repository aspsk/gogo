package counters

import "bytes"
import "testing"

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("12345"))
	if c != ByteCounter(5) {
		t.Errorf("expected c=5, but c=%d", int(c))
	}

}

func TestWordCounter(t *testing.T) {
	var c WordCounter
	c.Write([]byte("12345 \t123\t123     123  123"))
	if c != WordCounter(5) {
		t.Errorf("expected c=5, but c=%d", int(c))
	}
}

func TestLineCounter(t *testing.T) {
	var c LineCounter
	c.Write([]byte("12345\n234\n123\n123\n123"))
	if c != LineCounter(5) {
		t.Errorf("expected c=5, but c=%d", int(c))
	}
}

func TestCountingWriter(t *testing.T) {
	var buf bytes.Buffer
	ww, np := CountingWriter(&buf)
	ww.Write([]byte("foobar"))

	s := buf.String()
	if s != "foobar" {
		t.Errorf("expected buf.String()=foobar, but buf.String()=%s", s)
	}

	if *np != 6 {
		t.Errorf("expected *np=6, but *np=%d", *np)
	}
}
