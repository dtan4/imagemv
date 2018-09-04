package cli

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFlush(t *testing.T) {
	testcases := []struct {
		s    string
		want string
	}{
		{
			s:    "foobar",
			want: "foobar",
		},
		{
			s:    "",
			want: "",
		},
	}

	for _, tc := range testcases {
		got := &bytes.Buffer{}
		c := &concurrentWriter{
			w: bufio.NewWriter(got),
		}

		c.w.WriteString(tc.s)
		c.Flush()

		if got.String() != tc.want {
			t.Errorf("want: %q, got: %q", tc.want, got)
		}
	}
}

func TestWriteString(t *testing.T) {
	testcases := []struct {
		s    string
		want string
	}{
		{
			s:    "foobar",
			want: "foobar",
		},
		{
			s:    "",
			want: "",
		},
	}

	for _, tc := range testcases {
		got := &bytes.Buffer{}
		c := &concurrentWriter{
			w: bufio.NewWriter(got),
		}

		c.WriteString(tc.s)
		c.w.Flush()

		if got.String() != tc.want {
			t.Errorf("want: %q, got: %q", tc.want, got)
		}
	}
}

func BenchmarkBufferedWriter(b *testing.B) {
	c := &concurrentWriter{
		w: bufio.NewWriter(ioutil.Discard),
	}
	defer c.Flush()

	for i := 0; i < b.N; i++ {
		c.WriteString("foobarbaz")
	}
}

func BenchmarkStraightWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintln(ioutil.Discard, "foobarbaz")
	}
}
